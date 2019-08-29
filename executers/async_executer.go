package executers

import (
	"dogego/global"
	"dogego/modules"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func AsyncExecuter(ch *amqp.Channel, queue *amqp.Queue) {
	modules.RedisMQModule.Custome(
		global.TimeTaskQueueKey(),
		executeAsyncTask,
	)
}

func executeAsyncTask(message string) error {
	for _, item := range modules.TasksModule {
		l := strings.Split(message, "#$#")
		if item.Taskname == l[0] {
			var data interface{}

			if !modules.LockerModule.Lock(item.Taskname, 0) {
				return errors.New("Lock error.")
			}

			job := item.Job.(modules.AsyncTask)
			err := json.Unmarshal([]byte(l[1]), &data)

			if err != nil {
				modules.LockerModule.Free(item.Taskname)
				return err
			}

			from := time.Now().UnixNano()
			err = job(data)
			to := time.Now().UnixNano()

			if err != nil {
				log.Printf("%s Execute Error: %dms\n", item.Taskname, (to-from)/int64(time.Millisecond))
			} else {
				log.Printf("%s Execute Success: %dms\n", item.Taskname, (to-from)/int64(time.Millisecond))
			}

			modules.LockerModule.Free(item.Taskname)
		}
	}

	return nil
}
