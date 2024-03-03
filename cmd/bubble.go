package cmd

import (
	"fmt"

	"github.com/mdayat/sorta/internal"
	"github.com/spf13/cobra"
)

var (
	verbose     bool
	demonstrate bool
	bubbleCmd   = &cobra.Command{
		Use:   "bubble",
		Short: "Straightforward bubble sort description with an example",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(">>> What is Bubble Sort? <<<")
			fmt.Println(definition())

			fmt.Println(">>> Pros <<<")
			fmt.Println(pros())

			fmt.Println(">>> Cons <<<")
			fmt.Print(cons())

			if verbose {
				fmt.Print("\n>>> Implementation <<<\n")
				fmt.Print(implementation())
			}

			if demonstrate {
				fmt.Print("\n>>> Example <<<\n")

				oldNums := []int{2, 3, 1, 5, 4}
				newNums := internal.BubbleSort(oldNums)

				fmt.Println("Old List: ", oldNums)
				fmt.Println("New List: ", newNums)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(bubbleCmd)

	bubbleCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"a detailed description of bubble sort",
	)

	bubbleCmd.PersistentFlags().BoolVarP(
		&demonstrate,
		"demonstrate",
		"d",
		false,
		"demonstrate the bubble sort",
	)
}

func definition() string {
	definition := fmt.Sprintf(
		"Bubble sort is a simple sorting algorithm that repeatedly steps through the input list, \n%s",
		"element by element, comparing the current element with the one after it.\n",
	)
	return definition
}

func pros() string {
	pros := fmt.Sprintln("- Easy to write.")
	pros = pros + fmt.Sprintln("- Useful for a quick script when the amount of data is small.")
	return pros
}

func cons() string {
	cons := fmt.Sprintln("- One of the slowest sorting algorithms.")
	return cons
}

func implementation() string {
	impl := fmt.Sprintln("- Set swapping to true")
	impl = impl + fmt.Sprintln("- Set end to the (length of list - 1)")
	impl = impl + fmt.Sprintln("- While swapping is true")
	impl = impl + fmt.Sprintln("  - Set swapping to false")
	impl = impl + fmt.Sprintln("  - For i from 1st element to end")
	impl = impl + fmt.Sprintln("    - If the ith element of nums is greater than the (i+1)th element:")
	impl = impl + fmt.Sprintln("      - Swap the ith element and the (i+1)th element of nums")
	impl = impl + fmt.Sprintln("      - Set swapping to true")
	impl = impl + fmt.Sprintln("  - Reduce end by one")
	return impl
}
