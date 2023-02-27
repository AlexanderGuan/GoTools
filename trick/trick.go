/* Package trick implements some tricks in development */
package trick

import (
	"log"
	"time"
)

/* Record the time cost of this function */
func TestFunc() {
	defer func(startTime int64) {
		log.Printf("time cost: %v ms", (time.Now().UnixNano()-startTime)/1e6)
	}(time.Now().Unix())

	/* business logic below*/
}
