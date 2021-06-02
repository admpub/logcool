package cmd

import (
	"context"
	"fmt"
	"testing"
)

func Test_Command(t *testing.T) {
	Command(context.Background(), `
		{
    	"input": [
        	{
         	   "type": "stdin"
        	}
    	],
    	"filter": [
        	{
        	    "type": "zeus",
        	    "key": "foo",
        	    "value": "bar"
       	 	}
   	 	],
    	"output": [
        	{
            	"type": "stdout"
       	 	}
    	]
	}`)
}

func Test_Custom(t *testing.T) {
	Custom(context.Background(), "../templates/stdin2stdout.json")
}

func Test_LoadTemplates(t *testing.T) {
	LoadTemplates(context.Background())
}

func Test_Run(t *testing.T) {
	confs := Custom(context.Background(), "../templates/stdin2stdout.json")
	err := Run(confs)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_Help(t *testing.T) {
	Help()
}

func Test_Version(t *testing.T) {
	Version()
}
