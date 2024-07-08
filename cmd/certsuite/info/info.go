package info

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/test-network-function/cnf-certification-test/cnf-certification-test/identifiers"
	"github.com/test-network-function/cnf-certification-test/internal/cli"
	"github.com/test-network-function/test-network-function-claim/pkg/claim"
)

var (
	infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Displays information from the test catalog",
		RunE:  showInfo,
	}
)

func showInfo(cmd *cobra.Command, _ []string) error {
	testCaseFlag, _ := cmd.Flags().GetString("test-case")

	var testCase claim.TestCaseDescription
	for id := range identifiers.Catalog {
		if id.Id == testCaseFlag {
			testCase = identifiers.Catalog[id]
			break
		}
	}

	if testCase.Identifier.Id == "" {
		return fmt.Errorf("no test case found matching name %q", testCaseFlag)
	}

	const lineMaxWidth = 120

	// Test case identifier
	border := strings.Repeat("-", lineMaxWidth+4) //nolint:mnd // magic number
	fmt.Println(border)
	fmt.Printf("| %s |\n", cli.LineColor(cli.LineAlignCenter(testCase.Identifier.Id, lineMaxWidth), cli.Cyan))

	// Description
	border = strings.Repeat("-", lineMaxWidth+4) //nolint:mnd // magic number
	fmt.Println(border)
	fmt.Printf("| %s |\n", cli.LineColor(cli.LineAlignCenter("DESCRIPTION", lineMaxWidth), cli.Green))
	fmt.Println(border)
	for _, line := range cli.WrapLines(testCase.Description, lineMaxWidth) {
		fmt.Printf("| %s |\n", cli.LineAlignLeft(line, lineMaxWidth))
	}

	// Remediation
	fmt.Println(border)
	fmt.Printf("| %s |\n", cli.LineColor(cli.LineAlignCenter("REMEDIATION", lineMaxWidth), cli.Green))
	fmt.Println(border)
	for _, line := range cli.WrapLines(testCase.Remediation, lineMaxWidth) {
		fmt.Printf("| %s |\n", cli.LineAlignLeft(line, lineMaxWidth))
	}

	// Exceptions
	fmt.Println(border)
	fmt.Printf("| %s |\n", cli.LineColor(cli.LineAlignCenter("EXCEPTIONS", lineMaxWidth), cli.Green))
	fmt.Println(border)
	for _, line := range cli.WrapLines(testCase.ExceptionProcess, lineMaxWidth) {
		fmt.Printf("| %s |\n", cli.LineAlignLeft(line, lineMaxWidth))
	}

	// Best Practices reference
	fmt.Println(border)
	fmt.Printf("| %s |\n", cli.LineColor(cli.LineAlignCenter("BEST PRACTICES REFERENCE", lineMaxWidth), cli.Green))
	fmt.Println(border)
	for _, line := range cli.WrapLines(testCase.BestPracticeReference, lineMaxWidth) {
		fmt.Printf("| %s |\n", cli.LineAlignLeft(line, lineMaxWidth))
	}
	fmt.Println(border)

	return nil
}

func NewCommand() *cobra.Command {
	infoCmd.PersistentFlags().StringP("test-case", "t", "", "The test case to display information about")
	err := infoCmd.MarkPersistentFlagRequired("test-case")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not mark persistent flag \"test-case\" as required, err: %v", err)
		return nil
	}
	return infoCmd
}
