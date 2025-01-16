/*
Copyright © 2025 Dillon de Silva
*/
package cmd

import (
	"os"
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/go-git/go-git/v5/config"
	"github.com/teris-io/shortid"
	"os/exec"
	"cherrybomb/constants"
)

// getGitConfig returns the global git configuration
func getGitConfig() (*config.Config, error) {
	return config.LoadConfig(config.GlobalScope)
}

// Validate that at least one argument has been passed in
func getSourceBranchName(args *[]string) string {
	if len(*args) < 1 {
		fmt.Println("Please specify a source branch for cherrybomb to pick from")
		os.Exit(1)
	}

	return (*args)[0]
}

// Checkout a target branch and return the branch name
func checkoutTargetBranch() string {
	generatedBranchId, _ := shortid.Generate()
	targetBranchName := constants.AppName + "-" + generatedBranchId
	err := exec.Command("git", "checkout", "-b", targetBranchName, "--no-track", "origin/main").Run()
	if err != nil {
		log.Fatal("Failed to create target branch to bomb")
	}
	
	return targetBranchName
}

// Fetch from the upstream repository
func fetchUpstream(sourceBranchName *string) {
	err := exec.Command("git", "fetch", "upstream", *sourceBranchName).Run()
	if err != nil {
		log.Fatalf("Failed to fetch from upstream for branch %s: %s", *sourceBranchName, err)
	}
	log.Printf("%s-logger: Succeeding in fetching from source branch")
}

func getCommitHashes(authorEmail *string, sourceBranchName *string) []string {
	out, err := exec.Command("git", "log", "--left-right", "--graph", 
	"--cherry-pick", "--no-merges", "--author", *authorEmail, "--oneline",
	"--porcelain", "upstream/master..upstream/", sourceBranchName).CombinedOutput()
	if err != nil {
		log.Fatalf("Getting commit hashes %s: %s", *sourceBranchName, err)
	}
	log.Println(out)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cherrybomb",
	Short: "Utility for previewing and cherry-picking several commit hashes that ONLY you have authored, excluding merge commits",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Example of getting git config
		sourceBranchName := getSourceBranchName(&args)
		log.Printf("%s-logger: Running %s with source branch %s", 
		constants.AppName, constants.AppName, sourceBranchName)

		generatedBranchId, _ := shortid.Generate()
		targetBranchName := constants.AppName + "-" + generatedBranchId 

		log.Printf("%s-logger: Fetching from origin main\n", constants.AppName)
		err := exec.Command("git", "fetch", "origin", "main").Run()
		if err != nil {
			fmt.Printf("Error on fetching origin main. %s\n", err)
			log.Fatalf("%s-logger: %s\n", constants.AppName, err)
		}

		checkoutTargetBranch()
		fetchUpstream()
		getCommitHashes()
		

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


