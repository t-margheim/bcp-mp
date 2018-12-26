package canticles

import (
	"reflect"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want []Canticle
	}{
		{
			name: "December 24, 2018",
			date: time.Date(2018, 12, 24, 0, 0, 0, 0, time.UTC),
			want: []Canticle{
				{
					EnglishTitle: "The Song of Moses",
					LatinTitle:   "Cantemus Domino",
					Content:      "<p>I will sing to the Lord, for he is lofty and uplifted; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the horse and its rider has he hurled into the sea. <br/>The Lord is my strength and my refuge; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the Lord has become my Savior. <br/>This is my God and I will praise him, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the God of my people and I will exalt him. <br/>The Lord is a mighty warrior; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;Yahweh is his Name. <br/>The chariots of Pharaoh and his army has he hurled into the sea; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the finest of those who bear armor have been <br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;drowned in the Red Sea. <br/>The fathomless deep has overwhelmed them; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;they sank into the depths like a stone. <br/>Your right hand, O Lord, is glorious in might; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;your right hand, O Lord, has overthrown the enemy. <br/>Who can be compared with you, O Lord, among the gods? * <br/>&nbsp;&nbsp;&nbsp;&nbsp;who is like you, glorious in holiness, <br/>&nbsp;&nbsp;&nbsp;&nbsp;awesome in renown, and worker of wonders? <br/>You stretched forth your right hand; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the earth swallowed them up. <br/>With your constant love you led the people you redeemed; * <br/>&nbsp;&nbsp;&nbsp;&nbsp;with your might you brought them in safety to <br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;your holy dwelling. <br/>You will bring them in and plant them * <br/>&nbsp;&nbsp;&nbsp;&nbsp;on the mount of your possession, <br/>The resting-place you have made for yourself, O Lord, * <br/>&nbsp;&nbsp;&nbsp;&nbsp;the sanctuary, O Lord, that your hand has established. <br/>The Lord shall reign * <br/>&nbsp;&nbsp;&nbsp;&nbsp;for ever and for ever.</p><p>Glory to the Father, and to the Son, and to the Holy Spirit: * <br/>&nbsp;&nbsp;&nbsp;&nbsp;as it was in the beginning, is now, and will be for ever. Amen.</p>",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
