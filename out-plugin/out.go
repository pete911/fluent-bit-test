package main

// based on https://github.com/fluent/fluent-bit-go/tree/master/examples/out_multiinstance
import (
	"C"
	"fmt"
	"github.com/fluent/fluent-bit-go/output"
	"log"
	"os"
	"strings"
	"time"
	"unsafe"
)

var logFile *os.File

func init() {

	// directory "/tmp/fluent-bit-test" is mounted in values.yml
	f, err := os.Create("/tmp/fluent-bit-test/out")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	logFile = f
	log.SetFlags(0)
	log.SetOutput(logFile)
}

//export FLBPluginRegister
func FLBPluginRegister(def unsafe.Pointer) int {
	return output.FLBPluginRegister(def, "fluent-bit-test", "Test output plugin.")
}

//export FLBPluginInit
func FLBPluginInit(plugin unsafe.Pointer) int {
	// retrieve id value from fluent-bit-test plugin config
	id := output.FLBPluginConfigKey(plugin, "id")
	log.Printf("[fluent-bit-test] id = %q", id)
	// Set the context to point to any Go variable
	output.FLBPluginSetContext(plugin, id)

	return output.FLB_OK
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	log.Print("[fluent-bit-test] Flush called for unknown instance")
	return output.FLB_OK
}

//export FLBPluginFlushCtx
func FLBPluginFlushCtx(ctx, data unsafe.Pointer, length C.int, tag *C.char) int {
	// Type assert context back into the original type for the Go variable
	id := output.FLBPluginGetContext(ctx).(string)
	log.Printf("[fluent-bit-test] Flush called for id: %s", id)

	dec := output.NewDecoder(data, int(length))

	count := 0
	for {
		ret, ts, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		// this timestamp is coming from docker
		var timestamp time.Time
		switch t := ts.(type) {
		case output.FLBTime:
			timestamp = ts.(output.FLBTime).Time
		case uint64:
			timestamp = time.Unix(int64(t), 0)
		default:
			log.Println("time provided invalid, defaulting to now.")
			timestamp = time.Now()
		}

		var recordList []string
		for k, v := range record {
			recordList = append(recordList, fmt.Sprintf("    %s: %s\n", k, v))
		}

		// tag example
		// kube.var.log.containers.etcd-fluent-bit-test-control-plane_kube-system_etcd-9a20c0090877e231a66e9aa7018b0dc3dea77a0568512e7d85aa83b3cc22166e.log
		// removing last hash to make logs shorter
		tagParts := strings.Split(C.GoString(tag), ".")
		if len(tagParts) == 6 {
			containerParts := strings.Split(tagParts[4], "-")
			containerString := strings.Join(containerParts[:len(containerParts)-1], "-")
			tagParts[4] = fmt.Sprintf("%s...", containerString)
		}
		tagString := strings.Join(tagParts, ".")

		log.Printf("count: %d tag: %v\n  ts_record: %s\n  ts_output: %s\n  record:\n%+v\n\n",
			count, tagString, timestamp.Format(time.RFC3339Nano), time.Now().Format(time.RFC3339Nano), strings.Join(recordList, ""))
		count++
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	log.Print("[fluent-bit-test] Exit called for unknown instance")
	logFile.Close()
	return output.FLB_OK
}

//export FLBPluginExitCtx
func FLBPluginExitCtx(ctx unsafe.Pointer) int {
	return output.FLB_OK
}

func main() {
}
