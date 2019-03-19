package canticles

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		iterator int
		want     []Canticle
	}{
		{
			name:     "December 27, 2018",
			iterator: 360,
			want: []Canticle{
				{
					EnglishTitle: "The Song of Moses",
					LatinTitle:   "Cantemus Domino",
					Content:      "<p>I will sing to the Lord, for he is lofty and uplifted; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the horse and its rider has he hurled into the sea. <br/>The Lord is my strength and my refuge; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the Lord has become my Savior. <br/>This is my God and I will praise him, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the God of my people and I will exalt him. <br/>The Lord is a mighty warrior; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;Yahweh is his Name. <br/>The chariots of Pharaoh and his army has he hurled into the sea; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the finest of those who bear armor have been <br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;drowned in the Red Sea. <br/>The fathomless deep has overwhelmed them; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;they sank into the depths like a stone. <br/>Your right hand, O Lord, is glorious in might; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;your right hand, O Lord, has overthrown the enemy. <br/>Who can be compared with you, O Lord, among the gods? * <br/>&nbsp;&nbsp;&nbsp;&nbsp;who is like you, glorious in holiness, <br/>&nbsp;&nbsp;&nbsp;&nbsp;awesome in renown, and worker of wonders? <br/>You stretched forth your right hand; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the earth swallowed them up. <br/>With your constant love you led the people you redeemed; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;with your might you brought them in safety to <br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;your holy dwelling. <br/>You will bring them in and plant them * <br/>&nbsp;&nbsp;&nbsp;&nbsp;on the mount of your possession, <br/>The resting-place you have made for yourself, O Lord, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the sanctuary, O Lord, that your hand has established. <br/>The Lord shall reign * <br/>&nbsp;&nbsp;&nbsp;&nbsp;for ever and for ever.</p><p>Glory to the Father, and to the Son, and to the Holy Spirit: * <br/>&nbsp;&nbsp;&nbsp;&nbsp;as it was in the beginning, is now, and will be for ever. Amen.</p>",
				},
				{
					EnglishTitle: "A Song of Penitence",
					LatinTitle:   "Kyrie Pantokrator",
					Content:      "<p>O Lord and Ruler of the hosts of heaven, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;God of Abraham, Isaac, and Jacob, <br/>&nbsp;&nbsp;&nbsp;&nbsp;and of all their righteous offspring: <br/>You made the heavens and the earth, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;with all their vast array. </p> <p>All things quake with fear at your presence; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;they tremble because of your power. <br/>But your merciful promise is beyond all measure; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;it surpasses all that our minds can fathom. <br/>O Lord, you are full of compassion, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;long-suffering, and abounding in mercy. <br/>You hold back your hand; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;you do not punish as we deserve. <br/>In your great goodness, Lord, <br/>you have promised forgiveness to sinners, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;that they may repent of their sin and be saved. <br/>And now, O Lord, I bend the knee of my heart, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;and make my appeal, sure of your gracious goodness. <br/>I have sinned, O Lord, I have sinned, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;and I know my wickedness only too well. <br/>Therefore I make this prayer to you: * <br/>&nbsp;&nbsp;&nbsp;&nbsp;Forgive me, Lord, forgive me. <br/>Do not let me perish in my sin, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;nor condemn me to the depths of the earth. <br/>For you, O Lord, are the God of those who repent, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;and in me you will show forth your goodness. <br/>Unworthy as I am, you will save me, <br/>in accordance with your great mercy, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;and I will praise you without ceasing all the days of my life. <br/>For all the powers of heaven sing your praises, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;and yours is the glory to ages of ages. Amen.</p>",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.iterator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
