package skill

type TriggerType int

const (
	positive  TriggerType = iota + 1 //主动
	passive                          //被动
	addAttack                        //追击
	command                          //指挥
)

type level struct {
	Probability int   `json:"probability"`  //发动概率
	EffectValue []int `json:"effect_value"` //效果值
	EffectRound []int `json:"effect_round"` //效果持续回合数
}

type Conf struct {
	CfgId         int     `json:"cfgId"`
	Name          string  `json:"name"`
	Trigger       int     `json:"trigger"` //发起类型
	Target        int     `json:"target"`  //目标类型
	Des           string  `json:"des"`
	Limit         int     `json:"limit"`          //可以被武将装备上限
	Duration      int     `json:"duration"`       //技能持续时间，0为瞬时
	Arms          []int   `json:"arms"`           //可以装备的兵种
	IncludeEffect []int   `json:"include_effect"` //技能包括的效果
	Levels        []level `json:"levels"`
}

func (this *Conf) IsHitBefore() bool {
	return this.Trigger == int(positive) || this.Trigger == int(command)
}

func (this *Conf) IsHitAfter() bool {
	return this.Trigger == int(passive) || this.Trigger == int(addAttack)
}
