package credit

type CreditService interface {
	CheckCredit(string) bool
}

type creditService struct {
	Blacklist []string
}

func NewCreditService(blacklists []string) CreditService {
	return &creditService{blacklists}
}

// Return false if a person not pass background check
func (c *creditService) CheckCredit(firstName string) bool {
	for _, p := range c.Blacklist {
		if firstName == p {
			return false
		}
	}
	return true
}
