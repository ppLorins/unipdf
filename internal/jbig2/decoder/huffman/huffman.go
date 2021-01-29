//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package huffman ;import (_g "errors";_cb "fmt";_c "github.com/pplorins/unipdf/v3/internal/bitwise";_b "github.com/pplorins/unipdf/v3/internal/jbig2/internal";_e "math";_cf "strings";);func (_ge *EncodedTable )Decode (r _c .StreamReader )(int64 ,error ){return _ge ._dg .Decode (r )};func (_dgf *FixedSizeTable )InitTree (codeTable []*Code )error {_adca (codeTable );for _ ,_abf :=range codeTable {_bf :=_dgf ._da .append (_abf );if _bf !=nil {return _bf ;};};return nil ;};func NewEncodedTable (table BasicTabler )(*EncodedTable ,error ){_f :=&EncodedTable {_dg :&InternalNode {},BasicTabler :table };if _ec :=_f .parseTable ();_ec !=nil {return nil ,_ec ;};return _f ,nil ;};func (_fg *InternalNode )pad (_eeg *_cf .Builder ){for _gcc :=int32 (0);_gcc < _fg ._cec ;_gcc ++{_eeg .WriteString ("\u0020\u0020\u0020");};};type FixedSizeTable struct{_da *InternalNode };func NewCode (prefixLength ,rangeLength ,rangeLow int32 ,isLowerRange bool )*Code {return &Code {_egf :prefixLength ,_agb :rangeLength ,_gfc :rangeLow ,_fag :isLowerRange ,_gcg :-1};};func (_aa *InternalNode )Decode (r _c .StreamReader )(int64 ,error ){_ff ,_ddd :=r .ReadBit ();if _ddd !=nil {return 0,_ddd ;};if _ff ==0{return _aa ._cc .Decode (r );};return _aa ._ced .Decode (r );};var _ Node =&InternalNode {};func GetStandardTable (number int )(Tabler ,error ){if number <=0||number > len (_cae ){return nil ,_g .New ("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_cg :=_cae [number -1];if _cg ==nil {var _ceg error ;_cg ,_ceg =_db (_bfc [number -1]);if _ceg !=nil {return nil ,_ceg ;};_cae [number -1]=_cg ;};return _cg ,nil ;};func NewFixedSizeTable (codeTable []*Code )(*FixedSizeTable ,error ){_baa :=&FixedSizeTable {_da :&InternalNode {}};if _ab :=_baa .InitTree (codeTable );_ab !=nil {return nil ,_ab ;};return _baa ,nil ;};type ValueNode struct{_ggc int32 ;_bdg int32 ;_ag bool ;};func _ddg (_fc *Code )*OutOfBandNode {return &OutOfBandNode {}};type InternalNode struct{_cec int32 ;_cc Node ;_ced Node ;};type Tabler interface{Decode (_eegb _c .StreamReader )(int64 ,error );InitTree (_bef []*Code )error ;String ()string ;RootNode ()*InternalNode ;};func (_cedb *Code )String ()string {var _ea string ;if _cedb ._gcg !=-1{_ea =_dge (_cedb ._gcg ,_cedb ._egf );}else {_ea ="\u003f";};return _cb .Sprintf ("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_ea ,_cedb ._egf ,_cedb ._agb ,_cedb ._gfc );};type EncodedTable struct{BasicTabler ;_dg *InternalNode ;};func _db (_eee [][]int32 )(*StandardTable ,error ){var _cdc []*Code ;for _ad :=0;_ad < len (_eee );_ad ++{_bgg :=_eee [_ad ][0];_adc :=_eee [_ad ][1];_gcbc :=_eee [_ad ][2];var _bdgf bool ;if len (_eee [_ad ])> 3{_bdgf =true ;};_cdc =append (_cdc ,NewCode (_bgg ,_adc ,_gcbc ,_bdgf ));};_bcf :=&StandardTable {_cce :_eg (0)};if _bde :=_bcf .InitTree (_cdc );_bde !=nil {return nil ,_bde ;};return _bcf ,nil ;};func (_aag *InternalNode )String ()string {_cdf :=&_cf .Builder {};_cdf .WriteString ("\u000a");_aag .pad (_cdf );_cdf .WriteString ("\u0030\u003a\u0020");_cdf .WriteString (_aag ._cc .String ()+"\u000a");_aag .pad (_cdf );_cdf .WriteString ("\u0031\u003a\u0020");_cdf .WriteString (_aag ._ced .String ()+"\u000a");return _cdf .String ();};var _cae =make ([]Tabler ,len (_bfc ));type Code struct{_egf int32 ;_agb int32 ;_gfc int32 ;_fag bool ;_gcg int32 ;};var _ Node =&ValueNode {};var _ Node =&OutOfBandNode {};func (_ggg *StandardTable )InitTree (codeTable []*Code )error {_adca (codeTable );for _ ,_dfg :=range codeTable {if _bbg :=_ggg ._cce .append (_dfg );_bbg !=nil {return _bbg ;};};return nil ;};func (_ed *ValueNode )String ()string {return _cb .Sprintf ("\u0025\u0064\u002f%\u0064",_ed ._ggc ,_ed ._bdg );};func (_eeb *StandardTable )String ()string {return _eeb ._cce .String ()+"\u000a"};func _adca (_eac []*Code ){var _efg int32 ;for _ ,_gef :=range _eac {_efg =_bdee (_efg ,_gef ._egf );};_eegg :=make ([]int32 ,_efg +1);for _ ,_aac :=range _eac {_eegg [_aac ._egf ]++;};var _cbf int32 ;_egg :=make ([]int32 ,len (_eegg )+1);_eegg [0]=0;for _gac :=int32 (1);_gac <=int32 (len (_eegg ));_gac ++{_egg [_gac ]=(_egg [_gac -1]+(_eegg [_gac -1]))<<1;_cbf =_egg [_gac ];for _ ,_abg :=range _eac {if _abg ._egf ==_gac {_abg ._gcg =_cbf ;_cbf ++;};};};};func _bdee (_baag ,_eb int32 )int32 {if _baag > _eb {return _baag ;};return _eb ;};func (_be *InternalNode )append (_fdg *Code )(_bb error ){if _fdg ._egf ==0{return nil ;};_cdg :=_fdg ._egf -1-_be ._cec ;if _cdg < 0{return _g .New ("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064");};_fda :=(_fdg ._gcg >>uint (_cdg ))&0x1;if _cdg ==0{if _fdg ._agb ==-1{if _fda ==1{if _be ._ced !=nil {return _cb .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_fdg );};_be ._ced =_ddg (_fdg );}else {if _be ._cc !=nil {return _cb .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_fdg );};_be ._cc =_ddg (_fdg );};}else {if _fda ==1{if _be ._ced !=nil {return _cb .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_fdg );};_be ._ced =_cfg (_fdg );}else {if _be ._cc !=nil {return _cb .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_fdg );};_be ._cc =_cfg (_fdg );};};}else {if _fda ==1{if _be ._ced ==nil {_be ._ced =_eg (_be ._cec +1);};if _bb =_be ._ced .(*InternalNode ).append (_fdg );_bb !=nil {return _bb ;};}else {if _be ._cc ==nil {_be ._cc =_eg (_be ._cec +1);};if _bb =_be ._cc .(*InternalNode ).append (_fdg );_bb !=nil {return _bb ;};};};return nil ;};func (_gc *FixedSizeTable )RootNode ()*InternalNode {return _gc ._da };func _cfg (_gcb *Code )*ValueNode {return &ValueNode {_ggc :_gcb ._agb ,_bdg :_gcb ._gfc ,_ag :_gcb ._fag }};type StandardTable struct{_cce *InternalNode };func (_bac *ValueNode )Decode (r _c .StreamReader )(int64 ,error ){_fd ,_bg :=r .ReadBits (byte (_bac ._ggc ));if _bg !=nil {return 0,_bg ;};if _bac ._ag {_fd =-_fd ;};return int64 (_bac ._bdg )+int64 (_fd ),nil ;};func (_ecb *OutOfBandNode )Decode (r _c .StreamReader )(int64 ,error ){return 0,_b .ErrOOB };func (_fb *StandardTable )Decode (r _c .StreamReader )(int64 ,error ){return _fb ._cce .Decode (r )};func (_dd *EncodedTable )String ()string {return _dd ._dg .String ()+"\u000a"};func (_gg *EncodedTable )InitTree (codeTable []*Code )error {_adca (codeTable );for _ ,_cfd :=range codeTable {if _fa :=_gg ._dg .append (_cfd );_fa !=nil {return _fa ;};};return nil ;};func (_ggd *OutOfBandNode )String ()string {return _cb .Sprintf ("\u0025\u0030\u00364\u0062",int64 (_e .MaxInt64 ));};var _bfc =[][][]int32 {{{1,4,0},{2,8,16},{3,16,272},{3,32,65808}},{{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{6,32,75},{6,-1,0}},{{8,8,-256},{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{8,32,-257,999},{7,32,75},{6,-1,0}},{{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{5,32,76}},{{7,8,-255},{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{7,32,-256,999},{6,32,76}},{{5,10,-2048},{4,9,-1024},{4,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{2,7,0},{3,7,128},{3,8,256},{4,9,512},{4,10,1024},{6,32,-2049,999},{6,32,2048}},{{4,9,-1024},{3,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{4,5,0},{5,5,32},{5,6,64},{4,7,128},{3,8,256},{3,9,512},{3,10,1024},{5,32,-1025,999},{5,32,2048}},{{8,3,-15},{9,1,-7},{8,1,-5},{9,0,-3},{7,0,-2},{4,0,-1},{2,1,0},{5,0,2},{6,0,3},{3,4,4},{6,1,20},{4,4,22},{4,5,38},{5,6,70},{5,7,134},{6,7,262},{7,8,390},{6,10,646},{9,32,-16,999},{9,32,1670},{2,-1,0}},{{8,4,-31},{9,2,-15},{8,2,-11},{9,1,-7},{7,1,-5},{4,1,-3},{3,1,-1},{3,1,1},{5,1,3},{6,1,5},{3,5,7},{6,2,39},{4,5,43},{4,6,75},{5,7,139},{5,8,267},{6,8,523},{7,9,779},{6,11,1291},{9,32,-32,999},{9,32,3339},{2,-1,0}},{{7,4,-21},{8,0,-5},{7,0,-4},{5,0,-3},{2,2,-2},{5,0,2},{6,0,3},{7,0,4},{8,0,5},{2,6,6},{5,5,70},{6,5,102},{6,6,134},{6,7,198},{6,8,326},{6,9,582},{6,10,1094},{7,11,2118},{8,32,-22,999},{8,32,4166},{2,-1,0}},{{1,0,1},{2,1,2},{4,0,4},{4,1,5},{5,1,7},{5,2,9},{6,2,13},{7,2,17},{7,3,21},{7,4,29},{7,5,45},{7,6,77},{7,32,141}},{{1,0,1},{2,0,2},{3,1,3},{5,0,5},{5,1,6},{6,1,8},{7,0,10},{7,1,11},{7,2,13},{7,3,17},{7,4,25},{8,5,41},{8,32,73}},{{1,0,1},{3,0,2},{4,0,3},{5,0,4},{4,1,5},{3,3,7},{6,1,15},{6,2,17},{6,3,21},{6,4,29},{6,5,45},{7,6,77},{7,32,141}},{{3,0,-2},{3,0,-1},{1,0,0},{3,0,1},{3,0,2}},{{7,4,-24},{6,2,-8},{5,1,-4},{4,0,-2},{3,0,-1},{1,0,0},{3,0,1},{4,0,2},{5,1,3},{6,2,5},{7,4,9},{7,32,-25,999},{7,32,25}}};type BasicTabler interface{HtHigh ()int32 ;HtLow ()int32 ;StreamReader ()_c .StreamReader ;HtPS ()int32 ;HtRS ()int32 ;HtOOB ()int32 ;};type Node interface{Decode (_bd _c .StreamReader )(int64 ,error );String ()string ;};var _ Tabler =&EncodedTable {};func (_daa *StandardTable )RootNode ()*InternalNode {return _daa ._cce };func (_ce *EncodedTable )RootNode ()*InternalNode {return _ce ._dg };func (_ged *FixedSizeTable )Decode (r _c .StreamReader )(int64 ,error ){return _ged ._da .Decode (r )};func _eg (_ae int32 )*InternalNode {return &InternalNode {_cec :_ae }};func (_dae *FixedSizeTable )String ()string {return _dae ._da .String ()+"\u000a"};func _dge (_egd ,_cff int32 )string {var _bed int32 ;_gggc :=make ([]rune ,_cff );for _dad :=int32 (1);_dad <=_cff ;_dad ++{_bed =_egd >>uint (_cff -_dad )&1;if _bed !=0{_gggc [_dad -1]='1';}else {_gggc [_dad -1]='0';};};return string (_gggc );};type OutOfBandNode struct{};func (_ca *EncodedTable )parseTable ()error {var (_gf []*Code ;_fac ,_ee ,_dc int32 ;_a uint64 ;_cd error ;);_cfe :=_ca .StreamReader ();_ba :=_ca .HtLow ();for _ba < _ca .HtHigh (){_a ,_cd =_cfe .ReadBits (byte (_ca .HtPS ()));if _cd !=nil {return _cd ;};_fac =int32 (_a );_a ,_cd =_cfe .ReadBits (byte (_ca .HtRS ()));if _cd !=nil {return _cd ;};_ee =int32 (_a );_gf =append (_gf ,NewCode (_fac ,_ee ,_dc ,false ));_ba +=1<<uint (_ee );};_a ,_cd =_cfe .ReadBits (byte (_ca .HtPS ()));if _cd !=nil {return _cd ;};_fac =int32 (_a );_ee =32;_dc =_ca .HtLow ()-1;_gf =append (_gf ,NewCode (_fac ,_ee ,_dc ,true ));_a ,_cd =_cfe .ReadBits (byte (_ca .HtPS ()));if _cd !=nil {return _cd ;};_fac =int32 (_a );_ee =32;_dc =_ca .HtHigh ();_gf =append (_gf ,NewCode (_fac ,_ee ,_dc ,false ));if _ca .HtOOB ()==1{_a ,_cd =_cfe .ReadBits (byte (_ca .HtPS ()));if _cd !=nil {return _cd ;};_fac =int32 (_a );_gf =append (_gf ,NewCode (_fac ,-1,-1,false ));};if _cd =_ca .InitTree (_gf );_cd !=nil {return _cd ;};return nil ;};