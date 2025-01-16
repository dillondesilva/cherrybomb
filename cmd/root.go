/*
Copyright © 2025 Dillon de Silva
*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/go-git/go-git/v5/config"
	"github.com/teris-io/shortid"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"	
)

// getGitConfig returns the global git configuration
func getGitConfig() (*config.Config, error) {
	return config.LoadConfig(config.GlobalScope)
}

// Validate that at least one argument has been passed in
func verifyCherrybombArgs(args *[]string) {
	if len(*args) < 1 {
		fmt.Println("Please specify a source branch for cherrybomb to pick from")
		return
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cherrybomb",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Example of getting git config
		verifyCherrybombArgs(&args)
		generatedBranchId, _ := shortid.Generate()
		targetBranchName := "cherrybomb-" + generatedBranchId 
		cwd := os.Getwd()
		fmt.Println(targetBranchName)
		r, err := git.PlainOpen(cwd)
		cIter, err := r.Log()
		err = cIter.ForEach(func (c *object.Commit) error {
			fmt.Println(c)

			return nil
		}) 

		cfg, err := getGitConfig()
		if err != nil {
			fmt.Printf("Error loading git config: %v\n", err)
			return
		}

		// Get user email
		email := cfg.User.Email
		name := cfg.User.Name
		
		fmt.Printf("Git User Name: %s\n", name)
		fmt.Printf("Git User Email: %s\n", email)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cherrybomb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


