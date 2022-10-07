package godotenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestInitGodotenv(t *testing.T) {
	InitGodotenv()
	assert.EqualValues(t, "TEST", os.Getenv("RUN_ENV"))
}
