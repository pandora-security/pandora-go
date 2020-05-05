/**
 * Test and Example for Pandora
 **/
package pandora_test

import (
	"fmt"
	"github.com/danang-id/pandora-go"
	"github.com/danang-id/pandora-go/internal"
	"strings"
	"testing"
)

const ApplicationGUID = "00000000-0000-0000-0000-000000000000"
const TestWords = "Pandora Security Box"
var MinimumVersion = &internal.SemanticVersion{ Major: 0, Minor: 1, Patch: 0 }

func TestPandora_IsInstalled(t *testing.T) {
	t.Log("Creating Pandora instance")
	p := pandora.New(ApplicationGUID)
	t.Log("Check if Pandora is installed")
	isInstalled := p.IsInstalled()
	if !isInstalled {
		t.Error("Pandora is not detected on this system")
		fmt.Print("[TestPandora_IsInstalled] Pandora is not detected on this system")
		return
	}
	version, err := p.CheckVersion()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
		return
	}
	t.Logf("Pandora %s is detected on this system", version)
	fmt.Printf("[TestPandora_IsInstalled] Pandora %s is detected on this system\n", version)
}

func TestPandora_Cryptography(t *testing.T) {
	t.Log("Creating Pandora instance")
	p := pandora.New(ApplicationGUID)
	t.Logf("Setting minimum version to %s", MinimumVersion)
	p.SetMinimumVersion(MinimumVersion)
	t.Log()
	t.Logf("Encrypting \"%s\"", TestWords)
	app, cipher, err := p.Encrypt(TestWords)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
		return
	}
	t.Log("Response:")
	t.Logf(" => Application: %s [%s]", app.Name, app.Guid)
	t.Logf("                 by %s", app.Author)
	t.Logf(" => Cipher Text: %s", cipher)
	t.Log()
	t.Log("Decrypting back")
	app, plain, err := p.Decrypt(cipher)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
		return
	}
	t.Log("Response:")
	t.Logf(" => Application: %s [%s]", app.Name, app.Guid)
	t.Logf("                 by %s", app.Author)
	t.Logf(" => Plain Text : %s", plain)
	t.Log()
	var cryptoValidity string
	if strings.Compare(TestWords, string(plain)) == 0 {
		cryptoValidity = "VALID"
		t.Logf("Pandora cryptographic output is %s", cryptoValidity)
	} else {
		cryptoValidity = "NOT VALID"
		t.Errorf("Pandora cryptographic output is %s", cryptoValidity)
	}
	fmt.Printf("[TestPandora_Cryptography] Pandora cryptographic output is %s\n", cryptoValidity)
}
