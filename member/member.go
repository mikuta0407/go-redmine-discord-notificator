package member

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jszwec/csvutil"
)

type member struct {
	Name      string `csv:"DisplayName"`
	DiscordID string `csv:"DiscordID"`
}

func ConvertToDiscordID(toAddr string) (string, error) {
	localPartName := strings.Split(toAddr, "@")[0]

	memberList, err := getMemberList()
	if err != nil {
		return "", err
	}
	listLength := len(memberList)

	var discordID string

	for i, v := range memberList {
		if v.Name == localPartName {
			discordID = v.DiscordID
			break
		} else {
			if i == listLength-1 {
				return "", fmt.Errorf("%s", "Not Found")
			}
		}
	}
	return discordID, nil
}

func getMemberList() (memberList []member, err error) {
	dir, err := getDir()
	if err != nil {
		return
	}

	filePath := filepath.Join(dir, "MemberList.csv")
	b, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	if err = csvutil.Unmarshal(b, &memberList); err != nil {
		return
	}
	return
}

func getDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(exe, "/tmp/") {
		return os.Getwd()
	}
	return filepath.Dir(exe), nil
}
