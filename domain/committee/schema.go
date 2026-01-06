package committee

type RunCommitteeProcessInput struct {
	Question string
	Model    string
	Members  []string
	Opinion  bool
	Review   bool
	Stream   bool
}