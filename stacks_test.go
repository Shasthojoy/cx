package main

import (
	"flag"
	"github.com/cloud66/cli"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stacks list command", func() {
	Context("with a customer having three stacks in production", func() {

		var flagSet *flag.FlagSet

		BeforeEach(func() {
			//capture output and mock API endpoint
			StartCaptureStdout()
			MockApiGetCall("/stacks.json", 200, "./mocks/stacks/list.json")
			flagSet = flag.NewFlagSet("test", 0)
		})

		It("will show all three stacks running production", func() {
			// set the enviroment flag to production
			flagSet.String("environment", "production", "")

			// run context with the CLI
			context := cli.NewContext(nil, flagSet, nil)
			runStacks(context)

			// read stdout
			output := StopCaptureStdout()

			// check the actual output
			Expect(output[0]).To(HavePrefix("Awesome App1  production  Deployed successfully"))
			Expect(len(output) - 1).To(Equal(3))
		})

		It("will show all three stacks running production when no flag is specified", func() {
			// run context with the CLI
			context := cli.NewContext(nil, flagSet, nil)
			runStacks(context)

			// read stdout
			output := StopCaptureStdout()

			// check the actual output
			Expect(output[0]).To(HavePrefix("Awesome App1  production  Deployed successfully"))
			Expect(len(output) - 1).To(Equal(3))
		})

		It("will show all zero stacks running in development", func() {
			//  set the enviroment flag to development
			flagSet.String("environment", "development", "")

			// run context with the CLI
			context := cli.NewContext(nil, flagSet, nil)
			runStacks(context)

			// read stdout
			output := StopCaptureStdout()

			// check the actual output
			Expect(len(output) - 1).To(Equal(0))
		})

	})
})
