package modules

import (
	"fmt"
	"regexp"
	"strings"
)

func FindDAS(s string) (string, error) {
	var regex, err = regexp.Compile("DAS-[0-9]{1,4}")
	if err != nil {
		fmt.Println(err.Error())
		return "Error", err
	}

	v := regex.FindAllString(s, -1)
	if len(v) != 0 {
		return v[0], nil
	} else {
		return "", nil
	}
}

func ReplaceString(s string) string {
	repstr := strings.Replace(s, "Task", "", -1)
	repstr = strings.Replace(repstr, "Story", "", -1)
	repstr = strings.TrimSpace(repstr)

	return repstr
}

func FindAssignee(s string) (string, error) {
	var regex, err = regexp.Compile("Assignee: (.*)DAS")
	if err != nil {
		fmt.Println(err.Error())
		return "Error", err
	}

	v := regex.FindAllString(s, -1)
	if len(v) != 0 {
		get_v := v[0]
		get_v = strings.Replace(get_v, "DAS", "", -1)
		return get_v, nil
	} else {
		return "", nil
	}
}
