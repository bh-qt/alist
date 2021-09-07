package test

import (
	"fmt"
	"github.com/bh-qt/alist/conf"
	"github.com/bh-qt/alist/utils"
	"testing"
)

func TestStr(t *testing.T) {
	fmt.Println(".password-"[10:])
}

func TestWriteYml(t *testing.T) {
	utils.WriteToYml("../conf.yml", conf.Conf)
}
