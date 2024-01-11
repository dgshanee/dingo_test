package main

import (
	"dingo/components"
	"dingo/dom"
	"flag"
	"fmt"
	"os"
)

type UserPrompt struct{
	Prompts map[string]string;
}

var(
	user UserPrompt;
)

func buildCmd(){
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError);
	componentCmd := buildCmd.String("component", "", "Creates a component");
	dataCmd := buildCmd.String("data", "", "User provided data");
	IDCmd := buildCmd.String("id", "", "User provided ID");
	parentCmd := buildCmd.String("parent", "", "Parent of the created component");

	if len(os.Args) < 2{
		fmt.Println("Invalid command");
		return;
	}

	err := buildCmd.Parse(os.Args[2:]);
	if err != nil{
		fmt.Println("Error parsing commands");
		return;
	}

	user.Prompts = make(map[string]string);
	user.Prompts["component"] = *componentCmd;
	user.Prompts["data"] = *dataCmd;
	user.Prompts["id"] = *IDCmd;
	user.Prompts["parent"] = *parentCmd;
}

func main(){
	dom.LoadComponents();
	buildCmd();

	switch os.Args[1]{
	case "build":
		buildComponent();
	}
}

func buildComponent(){
	component, ok := user.Prompts["component"];
	if !ok{
		fmt.Println("No component specified");
		return;
	}
	switch component{
	case "text":
		var toMarshal components.TextComponent;
		components.PopulateStruct(&toMarshal, user.Prompts);
		dom.SaveComponent(toMarshal);
	default:
		panic("Invalid data type");
	}
	fmt.Println(user.Prompts)
}
