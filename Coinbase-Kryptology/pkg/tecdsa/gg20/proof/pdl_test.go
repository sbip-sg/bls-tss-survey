//
// Copyright Coinbase, Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package proof

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/stretchr/testify/require"

	tt "github.com/coinbase/kryptology/internal"
	crypto "github.com/coinbase/kryptology/pkg/core"
	"github.com/coinbase/kryptology/pkg/core/curves"
	paillier "github.com/coinbase/kryptology/pkg/paillier"
	"github.com/coinbase/kryptology/pkg/tecdsa/gg20/dealer"
)

func makeNewPaillierPublicKey(t *testing.T, n *big.Int) *paillier.PublicKey {
	t.Helper()
	publicKey, err := paillier.NewPubkey(n)
	require.NoError(t, err)
	return publicKey
}

func TestPdlProof(t *testing.T) {
	curve := btcec.S256()
	params := []*PdlProofParams{
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("10453436341595661308638945160744691461102262556435425363472041530426551492325701856213450266088099390020182300308521148275956201847403116901822304716109601")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("10963150595205830129496851796933410785873564590065350533653414494671353802134199060973898221372286631263311589006105573391294835324596497070909690030594529"),
				H1: tt.B10("443981490668637266429687390297201088749634563003174874283715634327157954734001973842075142869362000048454124614775810849589137797362291307830836093830666"),
				H2: tt.B10("552007421234219533984419815258072687324527963891381405280348833912665473111875648095851026421775527483650971365690299261791177376686729651798107556750986"),
			},
			C:       tt.B10("4664726001400863462852903463489684025203539312043087122435633780307370116923286498286011073893114770129124053266553951568210161226338715793626006742280222408787817580104395413260066838171514476835074847117030809020314721196250676941071342943510687061594910173132843733171341122020444569560262354325159459786"),
			ScalarR: tt.B10("6318180506937413445377662339737616688323403611418981231506634960717709610306639938046692771174821698278861338843502933136372971219816146088826194073514165"),
			ScalarX: tt.B10("8895158955508830352755492106542967678464513230702348453977370181840454571559"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("76094108851287611405923621794156813263013115318084984186541038825106228865281"),
				Y:     tt.B10("87411113406201178695670215615289769677595073686246590331802071626605210109743"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("8378869356347693656335563305413627471884674848153157571999293829483801014921"),
				Y:     tt.B10("80919981448100403067082012436084777199105270046135069091255381624902461382599"),
			},
		},
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("9831275438511106164565495760001688549814922646898355340010126520876442274445536039860446430205527149090890589584611773573077190211427866931337457355561521")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("9088113599264300599673127970260964968548172499135862549023291658586581399335917789448267260706627437027412917124323038435336668655448406702348643933116273"),
				H1: tt.B10("3164162665053366372408698072382121962273229838328732316301891025388141817123845974331611131026458919634038930316963586876688464406529778828717460449177018"),
				H2: tt.B10("3613863840368443617288882496207460760567029482693254218641355501439683045686086084409356221053813720993545273087552792708369718290403270154193540852884387"),
			},
			C:       tt.B10("22945662658826687122951524609605390162419018623016094527985190729136533122138237274890476792343375181731556056387967150127382752299222119038995116814564676335308780927745023238367652917410224228904840247506797913269791009050330669049148699546893290523696971497415351124890632011355109852928076331069242976047"),
			ScalarR: tt.B10("1456346843273107174185972100354877067846031294442486303631964113650903606327961364695375229133208501529871889372266089869069469359943307589956370077023244"),
			ScalarX: tt.B10("102510244919994967843708738855867526851165091005601348643599270378342008853277"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("86918276961810349294276103416548851884759982251107"),
				Y:     tt.B10("87194829221142880348582938487511785107150118762739500766654458540580527283772"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("12444758997853784844785736093813953015213016988344585020934721604959711243454"),
				Y:     tt.B10("41081235245407629494159229858676719985037195350871196150228771571530041697953"),
			},
		},
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("113704516018493703928170894572168412753803734325704651734551489944526965122231735113871864832368647984131374540170765278144857407900385316781049514749390883049421840883231007361862310299873591131872921717370699215293767892030993383657270453785323881620589052736661635699834501407565652575588677113465300661621")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("161396156400645198270315630211459947895162204967589557653128434044198146712972763644757188404750507699935523771400075723099627628687570310827922128423806613342745876762453195372210790970822267868149576881285953130288404538709228773231419065929483466982712730393834391664915100319075913286488970166494745830629"),
				H1: tt.B10("18618690214530067135386407378679435235583632518707852129591131476248349544570233992907989023204083131551404048200303628056048157230120466117628669672695188523977903126491955422473754106558050545355882619327085128114452064455293234627361539242255070012976950034765738988611881073014271066763434146268126349086"),
				H2: tt.B10("160470932600623889076430120482349695585923700506879108611495379697022889537530614221288183210040855444777056432425669107186536229798598233974027384740387767556408392083678883280773977871278647204040220814095879080051573171911139398073589016379401657293154896951331808925184453340679765040425291638017214591731"),
			},
			C:       tt.B10("4896957344081721172301761558878471817093957186883505910055484342858487410780344086015815560906271262760225544259435128094000760341287151463805517048421021266091498903664022813965524626993319004586190337532327563453628353189636257770661887076638254789941394875768438354317008447900448371463627310758169649159129081536916678375894293541030505690802142053594972076696842850896392543134018313744617812407507540361629152950683990975258614557530404228537967408981617289478956979359018250290787402363698811959464011054735419841399680079109795365937107391761037923382540359217859416039274471453061333934384874194680103175375"),
			ScalarR: tt.B10("71432491042795891026819073612920634365551967354604700070651069540194938209260355758675443987816023583431255584159007910956894102538667454292922436540633732978863886218440548882671875019295174919761937584083905373135768543476976716116826486377526415911236104646671272094245775861150495950556869225179514569498"),
			ScalarX: tt.B10("29086759249526442678185238367636452309232548742213292674117516677478809708340"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("86918276961810349294276103416548851884759982251107"),
				Y:     tt.B10("87194829221142880348582938487511785107150118762739500766654458540580527283772"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("35373660716694800516809670505588645945109927020371800385493928646782619107408"),
				Y:     tt.B10("93831376976173894628926808190406953951928006489246828107520931318493467317132"),
			},
		},
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("28389942534630501102482806139823952974111741690036242604042106439862886603688619267610452037222335475089537679295217755530571580626593285091009558531437945110816976195615021869354989046297137202661162264019525351295566584078998072196727905609637713692910458139555157767778141986605896692381096982746996945671616038237894860009730511432701694170671434630558138362258739746798822701440188845001261413779109525471954002454448797337162405307543709201259184397280458098930135776625337988818635617003108990612812385905691812906155358822153859313442892052035130822740985323920091265835286687462996096274931487719072325610177")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("24261699812894468219236852581962158941171546437127926412539821233148323141430665993557279634858451367749742151236664137456393535021164507932299356327658537600664706925498605224953136885130321375799440215429053090822764789891420452788542049702159277693971272237364975365363501274233385233018279005201887745841750070431152586842339528464629058292767675162224332893682681806584690404282462350130622536519356745051558971274796817643519111513887645688237884295068403256342873438983575942322642554660401209713049218828676032196860400247288848285743977850167317265395686177332360511824206274530781460468391577465417929720801"),
				H1: tt.B10("17657026647294223687309814976205454222798511294827726675315818303244321015282806894016959804822128590945638885512531773391303086945935616064555312507516711837306890638145180820603885391837297676456752054799643739324722052274736234615797408820276188727496193858254442729685055042577087166691843975492210965633375571812076567425573659274242387855645934243075788600138838255585659176549887756643115277634101423586780449892508521439191012182963941322615448360177625994713792119421184181502278161950043424758607370569862562496188656239924691950965423857119024515492868495595715338851989018661609371457796231769498336259942"),
				H2: tt.B10("23979964903511705755956518552835814554887788458101934011645738984289256815028997364100177641818315274085279710938880812881865338614102252653491843844747319033943655887178975358186429206512358407551818424689996649740298507695877144746406055658538582677662812963918275966732889580926799957382537995554279850611600698693796312231842307725821392566915849737659700063491348941930277735161335382722598575316473065306523191514953001778020716355562889584062417653273306819064551443565311996996530777393328559492161868722813868345420478163731017090977798561095759196393924713008544722136500375754011856471264672596599640334270"),
			},
			C:       tt.B10("312072118059655464479623752896370931287580695147558967175370385693010851904602632704190780873037853776949002099920910929454906746297000764022883262339810373530190503057275803212082490648752402411014710837921985963414326996430994445745103595174798048796872071058757782900936657543421311350819518047263143986520258140347967846766008789950384655455533378645886309519148467471177056605854541459719525084429218454218213528571466092270319677971542765590930957717305827486734849822346435125175072130311484595478227013437950067433362331207479003306966751040290017607880331032857536340225782113656911961520425583296868974183306576630007254668037105123599321810686598872455495139523140221304146923992336080310963429648303646680527705397398978349770824945433952717955829331451943599927661217135182414690477094098800503935476443441681845234327088836940113610943548236652767504866712759845009489949232476510935796276956643665275622378791333366356819315458176981133159815798879989329421090265495234364154899918069686655261893245095558625642862189807914671729788023888722841227953770641754100300708912019242483762170421213715045277109343120955558835679335136820251570631374672604143498438152345107896404355150794115568555566877217881851216716202636"),
			ScalarR: tt.B10("12900317669231902409891919243971118700371339340973561350358176269148021513243084734168439199346596288634953438481056037867044379527349847448380560153748923321339676070090389745139441984141990380824035678411732060617869487616828715411492104244485577512780240110853142143842539932038847998591353405830524617088326337057091796898754978001112967800124945178039660843489033360568972462506417939839595757327089284304802004150446809871851916545282806087722771656348015818813850000672740747971701015100804323897662661550498524602873193812403530774371979198086396784333450586513136371130153973494143884071046170273825540295768"),
			ScalarX: tt.B10("70007072993875354057337439998121157887778082743724503170686544130591021558682"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("86918276961810349294276103416548851884759982251107"),
				Y:     tt.B10("87194829221142880348582938487511785107150118762739500766654458540580527283772"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("69223975912566247908360773758405000717925830486569659950886484171255446147778"),
				Y:     tt.B10("32579403683948471182574476344537350186238230569150920925924612041995719512988"),
			},
		},
	}

	for _, pp := range params {
		proof, err := pp.Prove()
		if err != nil {
			t.Errorf("PdlProve failed: %v", err)
		}
		pv := &PdlVerifyParams{
			Curve:        curve,
			DealerParams: pp.DealerParams,
			Pk:           pp.Pk,
			PointX:       pp.PointX,
			PointR:       pp.PointR,
			C:            pp.C,
		}
		err = proof.Verify(pv)
		require.NoError(t, err)
	}
}

func TestPdlProofTampered(t *testing.T) {
	curve := btcec.S256()
	params := []*PdlProofParams{
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("10453436341595661308638945160744691461102262556435425363472041530426551492325701856213450266088099390020182300308521148275956201847403116901822304716109601")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("10963150595205830129496851796933410785873564590065350533653414494671353802134199060973898221372286631263311589006105573391294835324596497070909690030594529"),
				H1: tt.B10("443981490668637266429687390297201088749634563003174874283715634327157954734001973842075142869362000048454124614775810849589137797362291307830836093830666"),
				H2: tt.B10("552007421234219533984419815258072687324527963891381405280348833912665473111875648095851026421775527483650971365690299261791177376686729651798107556750986"),
			},
			C:       tt.B10("4664726001400863462852903463489684025203539312043087122435633780307370116923286498286011073893114770129124053266553951568210161226338715793626006742280222408787817580104395413260066838171514476835074847117030809020314721196250676941071342943510687061594910173132843733171341122020444569560262354325159459786"),
			ScalarR: tt.B10("6318180506937413445377662339737616688323403611418981231506634960717709610306639938046692771174821698278861338843502933136372971219816146088826194073514165"),
			ScalarX: tt.B10("8895158955508830352755492106542967678464513230702348453977370181840454571559"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("76094108851287611405923621794156813263013115318084984186541038825106228865281"),
				Y:     tt.B10("87411113406201178695670215615289769677595073686246590331802071626605210109743"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("8378869356347693656335563305413627471884674848153157571999293829483801014921"),
				Y:     tt.B10("80919981448100403067082012436084777199105270046135069091255381624902461382599"),
			},
		},
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("9831275438511106164565495760001688549814922646898355340010126520876442274445536039860446430205527149090890589584611773573077190211427866931337457355561521")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("9088113599264300599673127970260964968548172499135862549023291658586581399335917789448267260706627437027412917124323038435336668655448406702348643933116273"),
				H1: tt.B10("3164162665053366372408698072382121962273229838328732316301891025388141817123845974331611131026458919634038930316963586876688464406529778828717460449177018"),
				H2: tt.B10("3613863840368443617288882496207460760567029482693254218641355501439683045686086084409356221053813720993545273087552792708369718290403270154193540852884387"),
			},
			C:       tt.B10("22945662658826687122951524609605390162419018623016094527985190729136533122138237274890476792343375181731556056387967150127382752299222119038995116814564676335308780927745023238367652917410224228904840247506797913269791009050330669049148699546893290523696971497415351124890632011355109852928076331069242976047"),
			ScalarR: tt.B10("1456346843273107174185972100354877067846031294442486303631964113650903606327961364695375229133208501529871889372266089869069469359943307589956370077023244"),
			ScalarX: tt.B10("102510244919994967843708738855867526851165091005601348643599270378342008853277"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("86918276961810349294276103416548851884759982251107"),
				Y:     tt.B10("87194829221142880348582938487511785107150118762739500766654458540580527283772"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("12444758997853784844785736093813953015213016988344585020934721604959711243454"),
				Y:     tt.B10("41081235245407629494159229858676719985037195350871196150228771571530041697953"),
			},
		},
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("113704516018493703928170894572168412753803734325704651734551489944526965122231735113871864832368647984131374540170765278144857407900385316781049514749390883049421840883231007361862310299873591131872921717370699215293767892030993383657270453785323881620589052736661635699834501407565652575588677113465300661621")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("161396156400645198270315630211459947895162204967589557653128434044198146712972763644757188404750507699935523771400075723099627628687570310827922128423806613342745876762453195372210790970822267868149576881285953130288404538709228773231419065929483466982712730393834391664915100319075913286488970166494745830629"),
				H1: tt.B10("18618690214530067135386407378679435235583632518707852129591131476248349544570233992907989023204083131551404048200303628056048157230120466117628669672695188523977903126491955422473754106558050545355882619327085128114452064455293234627361539242255070012976950034765738988611881073014271066763434146268126349086"),
				H2: tt.B10("160470932600623889076430120482349695585923700506879108611495379697022889537530614221288183210040855444777056432425669107186536229798598233974027384740387767556408392083678883280773977871278647204040220814095879080051573171911139398073589016379401657293154896951331808925184453340679765040425291638017214591731"),
			},
			C:       tt.B10("4896957344081721172301761558878471817093957186883505910055484342858487410780344086015815560906271262760225544259435128094000760341287151463805517048421021266091498903664022813965524626993319004586190337532327563453628353189636257770661887076638254789941394875768438354317008447900448371463627310758169649159129081536916678375894293541030505690802142053594972076696842850896392543134018313744617812407507540361629152950683990975258614557530404228537967408981617289478956979359018250290787402363698811959464011054735419841399680079109795365937107391761037923382540359217859416039274471453061333934384874194680103175375"),
			ScalarR: tt.B10("71432491042795891026819073612920634365551967354604700070651069540194938209260355758675443987816023583431255584159007910956894102538667454292922436540633732978863886218440548882671875019295174919761937584083905373135768543476976716116826486377526415911236104646671272094245775861150495950556869225179514569498"),
			ScalarX: tt.B10("29086759249526442678185238367636452309232548742213292674117516677478809708340"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("86918276961810349294276103416548851884759982251107"),
				Y:     tt.B10("87194829221142880348582938487511785107150118762739500766654458540580527283772"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("35373660716694800516809670505588645945109927020371800385493928646782619107408"),
				Y:     tt.B10("93831376976173894628926808190406953951928006489246828107520931318493467317132"),
			},
		},
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("28389942534630501102482806139823952974111741690036242604042106439862886603688619267610452037222335475089537679295217755530571580626593285091009558531437945110816976195615021869354989046297137202661162264019525351295566584078998072196727905609637713692910458139555157767778141986605896692381096982746996945671616038237894860009730511432701694170671434630558138362258739746798822701440188845001261413779109525471954002454448797337162405307543709201259184397280458098930135776625337988818635617003108990612812385905691812906155358822153859313442892052035130822740985323920091265835286687462996096274931487719072325610177")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("24261699812894468219236852581962158941171546437127926412539821233148323141430665993557279634858451367749742151236664137456393535021164507932299356327658537600664706925498605224953136885130321375799440215429053090822764789891420452788542049702159277693971272237364975365363501274233385233018279005201887745841750070431152586842339528464629058292767675162224332893682681806584690404282462350130622536519356745051558971274796817643519111513887645688237884295068403256342873438983575942322642554660401209713049218828676032196860400247288848285743977850167317265395686177332360511824206274530781460468391577465417929720801"),
				H1: tt.B10("17657026647294223687309814976205454222798511294827726675315818303244321015282806894016959804822128590945638885512531773391303086945935616064555312507516711837306890638145180820603885391837297676456752054799643739324722052274736234615797408820276188727496193858254442729685055042577087166691843975492210965633375571812076567425573659274242387855645934243075788600138838255585659176549887756643115277634101423586780449892508521439191012182963941322615448360177625994713792119421184181502278161950043424758607370569862562496188656239924691950965423857119024515492868495595715338851989018661609371457796231769498336259942"),
				H2: tt.B10("23979964903511705755956518552835814554887788458101934011645738984289256815028997364100177641818315274085279710938880812881865338614102252653491843844747319033943655887178975358186429206512358407551818424689996649740298507695877144746406055658538582677662812963918275966732889580926799957382537995554279850611600698693796312231842307725821392566915849737659700063491348941930277735161335382722598575316473065306523191514953001778020716355562889584062417653273306819064551443565311996996530777393328559492161868722813868345420478163731017090977798561095759196393924713008544722136500375754011856471264672596599640334270"),
			},
			C:       tt.B10("312072118059655464479623752896370931287580695147558967175370385693010851904602632704190780873037853776949002099920910929454906746297000764022883262339810373530190503057275803212082490648752402411014710837921985963414326996430994445745103595174798048796872071058757782900936657543421311350819518047263143986520258140347967846766008789950384655455533378645886309519148467471177056605854541459719525084429218454218213528571466092270319677971542765590930957717305827486734849822346435125175072130311484595478227013437950067433362331207479003306966751040290017607880331032857536340225782113656911961520425583296868974183306576630007254668037105123599321810686598872455495139523140221304146923992336080310963429648303646680527705397398978349770824945433952717955829331451943599927661217135182414690477094098800503935476443441681845234327088836940113610943548236652767504866712759845009489949232476510935796276956643665275622378791333366356819315458176981133159815798879989329421090265495234364154899918069686655261893245095558625642862189807914671729788023888722841227953770641754100300708912019242483762170421213715045277109343120955558835679335136820251570631374672604143498438152345107896404355150794115568555566877217881851216716202636"),
			ScalarR: tt.B10("12900317669231902409891919243971118700371339340973561350358176269148021513243084734168439199346596288634953438481056037867044379527349847448380560153748923321339676070090389745139441984141990380824035678411732060617869487616828715411492104244485577512780240110853142143842539932038847998591353405830524617088326337057091796898754978001112967800124945178039660843489033360568972462506417939839595757327089284304802004150446809871851916545282806087722771656348015818813850000672740747971701015100804323897662661550498524602873193812403530774371979198086396784333450586513136371130153973494143884071046170273825540295768"),
			ScalarX: tt.B10("70007072993875354057337439998121157887778082743724503170686544130591021558682"),
			PointX: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("86918276961810349294276103416548851884759982251107"),
				Y:     tt.B10("87194829221142880348582938487511785107150118762739500766654458540580527283772"),
			},
			PointR: &curves.EcPoint{
				Curve: curve,
				X:     tt.B10("69223975912566247908360773758405000717925830486569659950886484171255446147778"),
				Y:     tt.B10("32579403683948471182574476344537350186238230569150920925924612041995719512988"),
			},
		},
	}

	for _, pp := range params {
		proof, err := pp.Prove()
		if err != nil {
			t.Errorf("ProvePDL failed: %v", err)
		}
		pv := &PdlVerifyParams{
			Curve:        curve,
			DealerParams: pp.DealerParams,
			Pk:           pp.Pk,
			PointX:       pp.PointX,
			PointR:       pp.PointR,
			C:            pp.C,
		}
		proof.z.Add(proof.z, crypto.One)
		if err := proof.Verify(pv); err == nil {
			t.Errorf("Expected PdlVerify to fail but succeeded.")
		}
		pv.C.Add(pv.C, crypto.One)
		if err := proof.Verify(pv); err == nil {
			t.Errorf("Expected PdlVerify to fail but succeeded.")
		}
		pv.C.Sub(pv.C, crypto.One)
		proof.s1.Add(proof.s1, crypto.One)
		if err := proof.Verify(pv); err == nil {
			t.Errorf("Expected PdlVerify to fail but succeeded.")
		}
		proof.s1.Sub(proof.s1, crypto.One)
		proof.e.Add(proof.e, crypto.One)
		if err := proof.Verify(pv); err == nil {
			t.Errorf("Expected PdlVerify to fail but succeeded.")
		}
	}
}

func TestPdlProofRandValues(t *testing.T) {
	curve := btcec.S256()
	n := tt.B10("10963150595205830129496851796933410785873564590065350533653414494671353802134199060973898221372286631263311589006105573391294835324596497070909690030594529")
	c, _ := crypto.Rand(n)
	r, _ := crypto.Rand(curve.Params().N)
	x, _ := crypto.Rand(curve.Params().N)
	pointX, _ := curves.NewScalarBaseMult(curve, x)
	pointR, _ := curves.NewScalarBaseMult(curve, r)
	params := []*PdlProofParams{
		{
			Curve: curve,
			Pk:    makeNewPaillierPublicKey(t, tt.B10("10453436341595661308638945160744691461102262556435425363472041530426551492325701856213450266088099390020182300308521148275956201847403116901822304716109601")),
			DealerParams: &dealer.ProofParams{
				N:  tt.B10("10963150595205830129496851796933410785873564590065350533653414494671353802134199060973898221372286631263311589006105573391294835324596497070909690030594529"),
				H1: tt.B10("443981490668637266429687390297201088749634563003174874283715634327157954734001973842075142869362000048454124614775810849589137797362291307830836093830666"),
				H2: tt.B10("552007421234219533984419815258072687324527963891381405280348833912665473111875648095851026421775527483650971365690299261791177376686729651798107556750986"),
			},
			C:       c,
			ScalarR: r,
			ScalarX: x,
			PointX:  pointX,
			PointR:  pointR,
		},
	}
	for _, pp := range params {
		proof, err := pp.Prove()
		if err != nil {
			t.Errorf("PdlProve failed: %v", err)
		}
		pv := &PdlVerifyParams{
			Curve:        curve,
			DealerParams: pp.DealerParams,
			Pk:           pp.Pk,
			PointX:       pp.PointX,
			PointR:       pp.PointR,
			C:            pp.C,
		}
		err = proof.Verify(pv)
		require.Error(t, err)
	}
}

func TestPdlProofInvalidParams(t *testing.T) {
	params := PdlProofParams{}
	if _, err := params.Prove(); err == nil {
		t.Errorf("PdlProve succeeded but should've failed.")
	}
	curve := btcec.S256()
	n := tt.B10("10963150595205830129496851796933410785873564590065350533653414494671353802134199060973898221372286631263311589006105573391294835324596497070909690030594529")
	c, _ := crypto.Rand(n)
	r, _ := crypto.Rand(curve.Params().N)
	x, _ := crypto.Rand(curve.Params().N)
	pointX, _ := curves.NewScalarBaseMult(curve, x)
	pointR, _ := curves.NewScalarBaseMult(curve, r)
	params = PdlProofParams{
		Curve: curve,
		Pk:    makeNewPaillierPublicKey(t, tt.B10("10453436341595661308638945160744691461102262556435425363472041530426551492325701856213450266088099390020182300308521148275956201847403116901822304716109601")),
		DealerParams: &dealer.ProofParams{
			N:  tt.B10("10963150595205830129496851796933410785873564590065350533653414494671353802134199060973898221372286631263311589006105573391294835324596497070909690030594529"),
			H1: tt.B10("443981490668637266429687390297201088749634563003174874283715634327157954734001973842075142869362000048454124614775810849589137797362291307830836093830666"),
			H2: tt.B10("552007421234219533984419815258072687324527963891381405280348833912665473111875648095851026421775527483650971365690299261791177376686729651798107556750986"),
		},
		C:       c,
		ScalarR: r,
		ScalarX: x,
		PointX:  pointX,
		PointR:  pointR,
	}
	proof, _ := params.Prove()
	pv := &PdlVerifyParams{}
	if err := proof.Verify(pv); err == nil {
		t.Errorf("PdlVerify succeeded but should've failed.")
	}
	pv.DealerParams = params.DealerParams
	if err := proof.Verify(pv); err == nil {
		t.Errorf("PdlVerify succeeded but should've failed.")
	}
	pv.Curve = params.Curve
	if err := proof.Verify(pv); err == nil {
		t.Errorf("PdlVerify succeeded but should've failed.")
	}
	pv.C = params.C
	if err := proof.Verify(pv); err == nil {
		t.Errorf("PdlVerify succeeded but should've failed.")
	}
	pv.PointR = params.PointR
	if err := proof.Verify(pv); err == nil {
		t.Errorf("PdlVerify succeeded but should've failed.")
	}
}

func TestMarshalJsonPdlProof(t *testing.T) {
	test := PdlProof{
		z:  bi(12),
		e:  bi(44),
		s:  bi(49),
		s1: bi(23),
		s2: bi(14),
	}

	testJSON, err := json.Marshal(test)
	require.NoError(t, err)
	require.NotNil(t, testJSON)

	unmarshaled := new(PdlProof)
	err = json.Unmarshal(testJSON, unmarshaled)
	require.NoError(t, err)

	require.Equal(t, test.z, unmarshaled.z)
	require.Equal(t, test.e, unmarshaled.e)
	require.Equal(t, test.s, unmarshaled.s)
	require.Equal(t, test.s1, unmarshaled.s1)
	require.Equal(t, test.s2, unmarshaled.s2)
}

func TestStaticJsonPdlProof(t *testing.T) {
	data := []byte(`{"Z":3666989406396256123443963142985320937042278632959430370804934299019701626404330930982847513018767734929278689417846246084124902745912477152349523658654602,"E":3115659271069276958682456068359865660696804370282039638783994681427792171677730593932305448178477745644721625256350270361466861753117055063398718073289695633908136229937717022554402500027743733983898897089676208130222278623864865543290659787169815752947365811271471278892825689889603560622879862341389394279,"S":77224884563607684241444406580257773576818651261464273214366432750766155500256211748871610535423213421390789952149378028436013126188776980935340867863856,"S1":739639048952431062486848243847124327050670851742270761995498294062920488651825727546170911992310856903527646746418983923284199570016337142789932710740405633070420584221244436733666126335959078297768150519780944901453061578676852438,"S2":7537969855339005139292750311314807059221711841400740408031810786514882275608436934507859710115709241956762803680015803927128791793177304895160503061938017049910995658419493513094195497328210909069904843769371376987211938248938850632714083738247153801422079209854600831288816219564593004239147913299330262096355556066672661530438723980217508272395618739481901662008216639924596404726636}`)
	expected := PdlProof{
		z:  tt.B10("3666989406396256123443963142985320937042278632959430370804934299019701626404330930982847513018767734929278689417846246084124902745912477152349523658654602"),
		e:  tt.B10("3115659271069276958682456068359865660696804370282039638783994681427792171677730593932305448178477745644721625256350270361466861753117055063398718073289695633908136229937717022554402500027743733983898897089676208130222278623864865543290659787169815752947365811271471278892825689889603560622879862341389394279"),
		s:  tt.B10("77224884563607684241444406580257773576818651261464273214366432750766155500256211748871610535423213421390789952149378028436013126188776980935340867863856"),
		s1: tt.B10("739639048952431062486848243847124327050670851742270761995498294062920488651825727546170911992310856903527646746418983923284199570016337142789932710740405633070420584221244436733666126335959078297768150519780944901453061578676852438"),
		s2: tt.B10("7537969855339005139292750311314807059221711841400740408031810786514882275608436934507859710115709241956762803680015803927128791793177304895160503061938017049910995658419493513094195497328210909069904843769371376987211938248938850632714083738247153801422079209854600831288816219564593004239147913299330262096355556066672661530438723980217508272395618739481901662008216639924596404726636"),
	}
	reconstructed := new(PdlProof)
	err := json.Unmarshal(data, reconstructed)
	require.NoError(t, err)
	require.Equal(t, expected.z, reconstructed.z)
	require.Equal(t, expected.e, reconstructed.e)
	require.Equal(t, expected.s, reconstructed.s)
	require.Equal(t, expected.s1, reconstructed.s1)
	require.Equal(t, expected.s2, reconstructed.s2)
}
