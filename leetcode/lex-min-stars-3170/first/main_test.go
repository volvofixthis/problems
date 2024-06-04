package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	i1 string
	o1 string
}{
	{"aaba*", "aab"},
	{"cb*aba*", "cab"},
	{"c*bba**", "b"},
	{"d*e*kdd***yd*y*kkk***y*ydyyy**yy***e*dde**edy***d***dd**k**k**yyyy****eddd****d*y*k*eyk*ee*yek*k****d**y**ykk**de*e***kyde*y*yk*y*k**kedkke****dd*y*k*e*d*e***k****y**ddkd***dykkd*e*ed****k****k*kkk***yd**k*ee**e*y*yk*dkee**k**ked*e*eey*kyydy*kk*****e*eekk*dkdyee***eyde*eee**kdekyy****dy***dy***dk*de**k***kdy*y*ed*k**e***kky**y**ye**d*****y*d**ek**d*****eke**e*ykkd*d*d***kdy*e*y*edek*y**yk*de***d***e***y***k***y*kyk*e*****d***y**ke**k*e*ey**d*dydyeky*dky*ykkk***dyykee*****y*e*d*keedy*ek**e******ykyyeee*k******dee*kky**ykddd******k*k*y*****e*e*dd*eke*k*k*e**kk*y****y*d*k**dydek**dkd******ddk*******k*e****de**e*e*e*kyy**kde*k*yky**kd****eyk**y*ykek*yddd**k***d**k*d*yyed**e****y**y*kkk*k*k*k*yy*e*ee*y*y*k***kk*******e*d*e*k*y*dkekkk*y*d**k*yde*****d**yk**d**d*k*k*y*e*k*e*k*k*d*k*ee**ky**eyeke*e*e**ky**e**y*y***e*e*k*k*yykkey****dek*kyy*k**d**k*kyeyk*ded**dy*ee*dy*ye**ykk*y*d*ekd*y****eydd*y**y**e*y*k*kkk**dd*y**yekd**ed***k*kk**k**eye*d*y*e*d*kd*kk*kydd****dyedd***deeyd************ke**k***k*kkd*k**ee*y*ky*yd***dy*d*ee*k*k****k**yy*yey**ed**eeky*ed****ey**dek***k***de*k*dd****k*d*e**keeee*y*k*k********ee**d*kye*k*d****yeekd****e*kedkd*k*eed*ed*e******dyy**kd*d*****ye***y*e*yy*e*ky**k**yk*y*d*ee*d****d**d**d*e**ed*k**kk**ey*kde*k*k*y*ky*k*d**eky***kd*d*y**k*ddk*yy*deke*ek***y*e*ydkdkkd**kdd*k*de**k*ek*k*dyy*e***k*******k**d*******ed*e*ky**d*kde*k*y*k*e**ey***yd*k*kky**yeyd****kk*e*deeddyyek***k**ykd*e*yd*y**kd*****e*eye*k*e*y****e*kk*e*y*ky**k**de*kd**d*k*kky****d*y*k**e**dkd**d****k******d***de*d**e***kkd**ddde***e*ke*ydedy**e**e***y*ey*e*****dd**y*ed**ky**de*d**d*y*k*y*dd**yye*dyke*dy*******dk*ed**yk***ee**y*y*ydye*k**d*y*y***yyye*k****dd*kd**eeeeeeeeeeeeaa**k", "keeeeeeeeeeeek"},
}

func BenchmarkCase1(b *testing.B) {
	clearStars(tests[0].i1)
}

func BenchmarkCase4(b *testing.B) {
	clearStars(tests[3].i1)
}

func TestFunc(t *testing.T) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, clearStars(test.i1), "should be equal")
		})
	}
}
