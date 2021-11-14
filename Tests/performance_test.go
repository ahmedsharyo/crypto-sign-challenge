package Tests

import (
	"runtime"
	"testing"
	"time"

	"github.com/ahmedsharyo/crypto-sign-challenge/Modules"
)

//tests SignPKCS1v15 aginst specific time limit
func Test_SignPKCS1v15_Time(t *testing.T) {

	//test case 1
	message := "Welcome to the Jungl"
	importedRSAPrivateKey := *Modules.LoadRSAPrivatePemKey(Modules.PrivateKeyPath)

	t0 := time.Now()
	Modules.SignPKCS1v15(message, importedRSAPrivateKey)

	t1 := time.Now()

	signingTime := t1.Sub(t0)

	if signingTime > 4*time.Millisecond {
		t.Errorf("Time limit exceeded on test case %d", 1) // to indicate test failed

	}

}

//-----------------------------------------------------

//tests SignPKCS1v15 aginst specific memory limit
func Test_SignPKCS1v15_MemUse(t *testing.T) {

	Modules.Save_keypair()

	//test case 1
	message := "Welcome to the Jungl"
	importedRSAPrivateKey := *Modules.LoadRSAPrivatePemKey(Modules.PrivateKeyPath)
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))
	var start, end runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&start)
	Modules.SignPKCS1v15(message, importedRSAPrivateKey)
	runtime.ReadMemStats(&end)
	alloc := end.TotalAlloc - start.TotalAlloc
	limit := uint64(40 * 1000)
	if alloc > limit {
		t.Error("memUse:", "allocated", alloc, "limit", limit)
	}
}
