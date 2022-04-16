package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(Hello())
	rootCmd.AddCommand(Start())
	rootCmd.AddCommand(Loader())

	rootCmd.Execute()
}

func Hello() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "Mundo", "flag para concatenar com Olá")

	return cmd
}

func Loader() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "load",
		Short: "just show load",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("process")
			loading()
		},
	}

	return cmd
}

func Start() *cobra.Command {

	var vars []string

	cmd := &cobra.Command{
		Use:   "start",
		Short: "just a test to stating cli process",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				name := StringPrompt("Nome: ")
				if name == "" {
					break
				}

				vars = append(vars, name)
			}
			fmt.Println(vars)
		},
	}

	return cmd

}

func loading() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf(".")
	}
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
