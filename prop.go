package goprop

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Prop struct {
	m map[string]string
}

func NewProp() *Prop {
	return &Prop{}
}

func (p *Prop) Get(key string) string {
	if key == "" {
		return ""
	}
	return p.m[key]
}

func (p *Prop) Read(propFileName string) error {
	file, err := os.OpenFile("prop.properties", os.O_RDONLY, 0660)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	p.m = make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		props := strings.SplitN(line, "=", 2)
		p.m[strings.TrimSpace(props[0])] = strings.TrimSpace(props[1])
		if err != nil || io.EOF == err {
			break
		}
	}
	return nil
}
