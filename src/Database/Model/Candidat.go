package model

type Candidat struct {
	Id               int
	Initial          string
	Email            string
	Phone            string
	Competence       string
	Experience       string
	Formation        string
	Firstname        string
	Lastname         string
	CreatedTime      string
	CandidateToOffer Offer
}
