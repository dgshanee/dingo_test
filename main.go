package main

import (
	"dingo/components"
	"dingo/dom"
	"flag"
	"fmt"
	"os"
	"strings"
)

type UserPrompt struct{
	Prompts map[string]string;
}

var(
	user UserPrompt;
)

func main(){
	dom.LoadComponents();

	buildCmd := flag.NewFlagSet("build", flag.ExitOnError);
	componentCmd := buildCmd.String("component", "", "Creates a component");
	getData := buildCmd.String("data", "", "User provided data");

	if len(os.Args) < 2{
		fmt.Println("Invalid command");
		return;
	}

	switch os.Args[1]{
	case "build":
		buildComponent(buildCmd, componentCmd, getData);
	}
}

func buildComponent(build *flag.FlagSet, component *string, data *string){
	user.Prompts = make(map[string]string);
	build.Parse(os.Args[2:]);
	for i:=2; i<len(os.Args); i+=2{
		if i+1<len(os.Args){
			key := strings.TrimPrefix(os.Args[i], "--");
			user.Prompts[key] = os.Args[i+1];
		}
	}

	switch *component{
	case "text":
		var toMarshal components.TextComponent;
		components.PopulateStruct(&toMarshal, user.Prompts);
		dom.SaveComponent(toMarshal);
	default:
		panic("Invalid data type");
	}
	fmt.Println(user.Prompts)
}
