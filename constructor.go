package blockrand

import "github.com/tyeolrik/rngset"

type participant struct {
	userID    string
	userInput uint64
}

type block struct {
	realSeed uint64

	participants []participant
}

type sr__mt19937_64__well19937a struct {
	// Information
	_blockSize uint16 // n개의 블록이 모이면 output을 만들 수 있다.

	state  uint64  // 현재 단계
	blocks []block // n개의 블록이 모이면 output을 만들 수 있음.
}

func NewSR(blockSize uint16) (ret sr__mt19937_64__well19937a) {
	ret = sr__mt19937_64__well19937a{
		_blockSize: blockSize,
		state:      0,
	}
	ret.blocks = make([]block, blockSize)
	for i := range ret.blocks {
		ret.blocks[i].realSeed = 0
		ret.blocks[i].participants = make([]participant, 0)
	}
	return
}

func (r *sr__mt19937_64__well19937a) Participate(userID string, userInput uint64) {
	r.blocks[r.state].participants = append(r.blocks[r.state].participants, participant{userID: userID, userInput: userInput})
}

func (r *sr__mt19937_64__well19937a) Mining() {
	// Calculate realSeed
	// MT19937_64 needs only 1 seed. So, all participant input should be XORed
	for i := range r.blocks[r.state].participants {
		r.blocks[r.state].realSeed = r.blocks[r.state].realSeed ^ r.blocks[r.state].participants[i].userInput
	}
	mt19937_64 := rngset.NewMT19937_64(r.blocks[r.state].realSeed)
	r.blocks[r.state].realSeed = mt19937_64.NextUint64()

	// Go Next Block State
	r.state++
}
