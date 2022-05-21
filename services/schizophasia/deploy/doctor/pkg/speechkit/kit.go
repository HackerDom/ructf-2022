package speechkit

import (
	"cpl.li/go/markov"
	"strings"
)

const text = "Cybersecurity is the practice of protecting critical systems and sensitive information from digital attacks Also known as information technology security cybersecurity measures are designed to combat threats against networked systems and applications whether those threats originate from inside or outside of an organization In 2020 the average cost of a data breach was 3 USD million globally and USD 8 million in the United States These costs include the expenses of discovering and responding to the breach the cost of downtime and lost revenue and the long-term reputational damage to a business and its brand Cybercriminals target customers personally identifiable information names addresses national identification numbers for example Social Security number in the US fiscal codes in Italy and credit card information and then sell these records in underground digital marketplaces Compromised PII often leads to a loss of customer trust the imposition of regulatory fines, and even legal action Security system complexity created by disparate technologies and a lack of in-house expertise can amplify these costs But organizations with a comprehensive cybersecurity strategy governed by best practices and automated using advanced analytics artificial intelligence and machine learning can fight cyberthreats more effectively and reduce the lifecycle and impact of breaches when they occur"
const maxWords = 100 // max words to generate (default 100)
const pairSize = 2   // size of a word pair (default 2)
var chain *markov.Chain

func Generate() string {
	b := chain.NewBuilder(nil) // create builder on top of chain

	b.Generate(maxWords - pairSize) // generate new words
	return b.String()
}

func init() {
	// handle flags
	chain = markov.NewChain(pairSize) // create markov chain

	// give data as sequence to chain model
	chain.Add(strings.Fields(text))
}
