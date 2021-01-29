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

package classer ;import (_g "github.com/pplorins/unipdf/v3/common";_a "github.com/pplorins/unipdf/v3/internal/jbig2/basic";_fc "github.com/pplorins/unipdf/v3/internal/jbig2/bitmap";_c "github.com/pplorins/unipdf/v3/internal/jbig2/errors";_e "image";_b "math";);type similarTemplatesFinder struct{Classer *Classer ;Width int ;Height int ;Index int ;CurrentNumbers []int ;N int ;};func _bdfcf (_ege *Classer ,_faaa *_fc .Bitmap )*similarTemplatesFinder {return &similarTemplatesFinder {Width :_faaa .Width ,Height :_faaa .Height ,Classer :_ege };};const (RankHaus Method =iota ;Correlation ;);func (_ga *Classer )addPageComponents (_cd *_fc .Bitmap ,_bea *_fc .Boxes ,_dc *_fc .Bitmaps ,_bf int ,_gf Method )error {const _eb ="\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s";if _cd ==nil {return _c .Error (_eb ,"\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065");};if _bea ==nil ||_dc ==nil ||len (*_bea )==0{_g .Log .Trace ("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064",_cd );return nil ;};var _bd error ;switch _gf {case RankHaus :_bd =_ga .classifyRankHaus (_bea ,_dc ,_bf );case Correlation :_bd =_ga .classifyCorrelation (_bea ,_dc ,_bf );default:_g .Log .Debug ("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027",_gf );return _c .Error (_eb ,"\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064");};if _bd !=nil {return _c .Wrap (_bd ,_eb ,"");};if _bd =_ga .getULCorners (_cd ,_bea );_bd !=nil {return _c .Wrap (_bd ,_eb ,"");};_bc :=len (*_bea );_ga .BaseIndex +=_bc ;if _bd =_ga .ComponentsNumber .Add (_bc );_bd !=nil {return _c .Wrap (_bd ,_eb ,"");};return nil ;};var TwoByTwoWalk =[]int {0,0,0,1,-1,0,0,-1,1,0,-1,1,1,1,-1,-1,1,-1,0,-2,2,0,0,2,-2,0,-1,-2,1,-2,2,-1,2,1,1,2,-1,2,-2,1,-2,-1,-2,-2,2,-2,2,2,-2,2};func _bga (_da *_fc .Bitmap ,_db ,_fbb ,_eba ,_edb int ,_cb *_fc .Bitmap )(_aae _e .Point ,_bgg error ){const _bec ="\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg";if _da ==nil {return _aae ,_c .Error (_bec ,"\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};if _cb ==nil {return _aae ,_c .Error (_bec ,"t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064");};_afd ,_fae :=_cb .Width ,_cb .Height ;_gee ,_bde :=_db -_eba -JbAddedPixels ,_fbb -_edb -JbAddedPixels ;_g .Log .Trace ("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027",_db ,_fbb ,_afd ,_fae ,_gee ,_bde );_gbd ,_bgg :=_fc .Rect (_gee ,_bde ,_afd ,_fae );if _bgg !=nil {return _aae ,_c .Wrap (_bgg ,_bec ,"");};_cc ,_ ,_bgg :=_da .ClipRectangle (_gbd );if _bgg !=nil {_g .Log .Error ("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076",_gbd );return _aae ,_c .Wrap (_bgg ,_bec ,"");};_dg :=_fc .New (_cc .Width ,_cc .Height );_ag :=_b .MaxInt32 ;var _dbf ,_fe ,_cfe ,_cbd ,_aga int ;for _dbf =-1;_dbf <=1;_dbf ++{for _fe =-1;_fe <=1;_fe ++{if _ ,_bgg =_fc .Copy (_dg ,_cc );_bgg !=nil {return _aae ,_c .Wrap (_bgg ,_bec ,"");};if _bgg =_dg .RasterOperation (_fe ,_dbf ,_afd ,_fae ,_fc .PixSrcXorDst ,_cb ,0,0);_bgg !=nil {return _aae ,_c .Wrap (_bgg ,_bec ,"");};_cfe =_dg .CountPixels ();if _cfe < _ag {_cbd =_fe ;_aga =_dbf ;_ag =_cfe ;};};};_aae .X =_cbd ;_aae .Y =_aga ;return _aae ,nil ;};type Method int ;const (MaxDiffWidth =2;MaxDiffHeight =2;);type Settings struct{MaxCompWidth int ;MaxCompHeight int ;SizeHaus int ;RankHaus float64 ;Thresh float64 ;WeightFactor float64 ;KeepClassInstances bool ;Components _fc .Component ;Method Method ;};func (_eef *Classer )verifyMethod (_gff Method )error {if _gff !=RankHaus &&_gff !=Correlation {return _c .Error ("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064","\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064");};return nil ;};func DefaultSettings ()Settings {_afdg :=&Settings {};_afdg .SetDefault ();return *_afdg };func (_gb *Classer )AddPage (inputPage *_fc .Bitmap ,pageNumber int ,method Method )(_ca error ){const _be ="\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065";_gb .Widths [pageNumber ]=inputPage .Width ;_gb .Heights [pageNumber ]=inputPage .Height ;if _ca =_gb .verifyMethod (method );_ca !=nil {return _c .Wrap (_ca ,_be ,"");};_ae ,_ecf ,_ca :=inputPage .GetComponents (_gb .Settings .Components ,_gb .Settings .MaxCompWidth ,_gb .Settings .MaxCompHeight );if _ca !=nil {return _c .Wrap (_ca ,_be ,"");};_g .Log .Debug ("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076",_ae );if _ca =_gb .addPageComponents (inputPage ,_ecf ,_ae ,pageNumber ,method );_ca !=nil {return _c .Wrap (_ca ,_be ,"");};return nil ;};func (_cfd *Classer )getULCorners (_dcb *_fc .Bitmap ,_ee *_fc .Boxes )error {const _ad ="\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _dcb ==nil {return _c .Error (_ad ,"\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070");};if _ee ==nil {return _c .Error (_ad ,"\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};if _cfd .PtaUL ==nil {_cfd .PtaUL =&_fc .Points {};};_gad :=len (*_ee );var (_ed ,_caf ,_bbd ,_bce int ;_aab ,_gbe ,_gfe ,_bg float32 ;_fd error ;_gbeb *_e .Rectangle ;_fcf *_fc .Bitmap ;_ge _e .Point ;);for _cag :=0;_cag < _gad ;_cag ++{_ed =_cfd .BaseIndex +_cag ;if _aab ,_gbe ,_fd =_cfd .CentroidPoints .GetGeometry (_ed );_fd !=nil {return _c .Wrap (_fd ,_ad ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073");};if _caf ,_fd =_cfd .ClassIDs .Get (_ed );_fd !=nil {return _c .Wrap (_fd ,_ad ,"\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074");};if _gfe ,_bg ,_fd =_cfd .CentroidPointsTemplates .GetGeometry (_caf );_fd !=nil {return _c .Wrap (_fd ,_ad ,"\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_fa :=_gfe -_aab ;_geb :=_bg -_gbe ;if _fa >=0{_bbd =int (_fa +0.5);}else {_bbd =int (_fa -0.5);};if _geb >=0{_bce =int (_geb +0.5);}else {_bce =int (_geb -0.5);};if _gbeb ,_fd =_ee .Get (_cag );_fd !=nil {return _c .Wrap (_fd ,_ad ,"");};_bca ,_gg :=_gbeb .Min .X ,_gbeb .Min .Y ;_fcf ,_fd =_cfd .UndilatedTemplates .GetBitmap (_caf );if _fd !=nil {return _c .Wrap (_fd ,_ad ,"\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029");};_ge ,_fd =_bga (_dcb ,_bca ,_gg ,_bbd ,_bce ,_fcf );if _fd !=nil {return _c .Wrap (_fd ,_ad ,"");};_cfd .PtaUL .AddPoint (float32 (_bca -_bbd +_ge .X ),float32 (_gg -_bce +_ge .Y ));};return nil ;};func (_bef *Classer )classifyCorrelation (_gbdd *_fc .Boxes ,_bdfc *_fc .Bitmaps ,_gd int )error {const _fba ="\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e";if _gbdd ==nil {return _c .Error (_fba ,"\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064");};if _bdfc ==nil {return _c .Error (_fba ,"\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");};_fcd :=len (_bdfc .Values );if _fcd ==0{_g .Log .Debug ("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079");return nil ;};var (_gfg ,_gbc *_fc .Bitmap ;_gdb error ;);_fca :=&_fc .Bitmaps {Values :make ([]*_fc .Bitmap ,_fcd )};for _aabf ,_df :=range _bdfc .Values {_gbc ,_gdb =_df .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);if _gdb !=nil {return _c .Wrap (_gdb ,_fba ,"");};_fca .Values [_aabf ]=_gbc ;};_ea :=_bef .FgTemplates ;_gebd :=_fc .MakePixelSumTab8 ();_fab :=_fc .MakePixelCentroidTab8 ();_ffa :=make ([]int ,_fcd );_gcc :=make ([][]int ,_fcd );_ffb :=_fc .Points (make ([]_fc .Point ,_fcd ));_fgg :=&_ffb ;var (_ggf ,_fde int ;_ada ,_dfg ,_fac int ;_bgd ,_cdg int ;_eeg byte ;);for _aac ,_fggg :=range _fca .Values {_gcc [_aac ]=make ([]int ,_fggg .Height );_ggf =0;_fde =0;_dfg =(_fggg .Height -1)*_fggg .RowStride ;_ada =0;for _cdg =_fggg .Height -1;_cdg >=0;_cdg ,_dfg =_cdg -1,_dfg -_fggg .RowStride {_gcc [_aac ][_cdg ]=_ada ;_fac =0;for _bgd =0;_bgd < _fggg .RowStride ;_bgd ++{_eeg =_fggg .Data [_dfg +_bgd ];_fac +=_gebd [_eeg ];_ggf +=_fab [_eeg ]+_bgd *8*_gebd [_eeg ];};_ada +=_fac ;_fde +=_fac *_cdg ;};_ffa [_aac ]=_ada ;if _ada > 0{(*_fgg )[_aac ]=_fc .Point {X :float32 (_ggf )/float32 (_ada ),Y :float32 (_fde )/float32 (_ada )};}else {(*_fgg )[_aac ]=_fc .Point {X :float32 (_fggg .Width )/float32 (2),Y :float32 (_fggg .Height )/float32 (2)};};};if _gdb =_bef .CentroidPoints .Add (_fgg );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064");};var (_eaf ,_aad ,_gaa int ;_eec float64 ;_dba ,_eg ,_dd ,_bbe float32 ;_beb ,_eac _fc .Point ;_afa bool ;_gbee *similarTemplatesFinder ;_cbf int ;_agb *_fc .Bitmap ;_cdd *_e .Rectangle ;_abc *_fc .Bitmaps ;);for _cbf ,_gbc =range _fca .Values {_aad =_ffa [_cbf ];if _dba ,_eg ,_gdb =_fgg .GetGeometry (_cbf );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0070t\u0061\u0020\u002d\u0020\u0069");};_afa =false ;_gbg :=len (_bef .UndilatedTemplates .Values );_gbee =_bdfcf (_bef ,_gbc );for _egg :=_gbee .Next ();_egg > -1;{if _agb ,_gdb =_bef .UndilatedTemplates .GetBitmap (_egg );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032");};if _gaa ,_gdb =_ea .GetInt (_egg );_gdb !=nil {_g .Log .Trace ("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_gdb );};if _dd ,_bbe ,_gdb =_bef .CentroidPointsTemplates .GetGeometry (_egg );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020");};if _bef .Settings .WeightFactor > 0.0{if _eaf ,_gdb =_bef .TemplateAreas .Get (_egg );_gdb !=nil {_g .Log .Trace ("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076",_gdb );};_eec =_bef .Settings .Thresh +(1.0-_bef .Settings .Thresh )*_bef .Settings .WeightFactor *float64 (_gaa )/float64 (_eaf );}else {_eec =_bef .Settings .Thresh ;};_bgf ,_edf :=_fc .CorrelationScoreThresholded (_gbc ,_agb ,_aad ,_gaa ,_beb .X -_eac .X ,_beb .Y -_eac .Y ,MaxDiffWidth ,MaxDiffHeight ,_gebd ,_gcc [_cbf ],float32 (_eec ));if _edf !=nil {return _c .Wrap (_edf ,_fba ,"");};if _ab {var (_dfa ,_ecd float64 ;_dbc ,_cda int ;);_dfa ,_edf =_fc .CorrelationScore (_gbc ,_agb ,_aad ,_gaa ,_dba -_dd ,_eg -_bbe ,MaxDiffWidth ,MaxDiffHeight ,_gebd );if _edf !=nil {return _c .Wrap (_edf ,_fba ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");};_ecd ,_edf =_fc .CorrelationScoreSimple (_gbc ,_agb ,_aad ,_gaa ,_dba -_dd ,_eg -_bbe ,MaxDiffWidth ,MaxDiffHeight ,_gebd );if _edf !=nil {return _c .Wrap (_edf ,_fba ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");};_dbc =int (_b .Sqrt (_dfa *float64 (_aad )*float64 (_gaa )));_cda =int (_b .Sqrt (_ecd *float64 (_aad )*float64 (_gaa )));if (_dfa >=_eec )!=(_ecd >=_eec ){return _c .Errorf (_fba ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066",_dbc ,_dfa ,_dfa >=_eec ,_cda ,_ecd ,_ecd >=_eec ,_dfa -_ecd );};if _dfa >=_eec !=_bgf {return _c .Errorf (_fba ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076",_dfa ,_dfa *float64 (_aad )*float64 (_gaa ),_dbc ,_eec ,float32 (_eec )*float32 (_aad )*float32 (_gaa ),_bgf );};};if _bgf {_afa =true ;if _edf =_bef .ClassIDs .Add (_egg );_edf !=nil {return _c .Wrap (_edf ,_fba ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");};if _edf =_bef .ComponentPageNumbers .Add (_gd );_edf !=nil {return _c .Wrap (_edf ,_fba ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");};if _bef .Settings .KeepClassInstances {if _gfg ,_edf =_bdfc .GetBitmap (_cbf );_edf !=nil {return _c .Wrap (_edf ,_fba ,"\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069");};if _abc ,_edf =_bef .ClassInstances .GetBitmaps (_egg );_edf !=nil {return _c .Wrap (_edf ,_fba ,"K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073");};_abc .AddBitmap (_gfg );if _cdd ,_edf =_gbdd .Get (_cbf );_edf !=nil {return _c .Wrap (_edf ,_fba ,"\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073");};_abc .AddBox (_cdd );};break ;};};if !_afa {if _gdb =_bef .ClassIDs .Add (_gbg );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _gdb =_bef .ComponentPageNumbers .Add (_gd );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_abc =&_fc .Bitmaps {};if _gfg ,_gdb =_bdfc .GetBitmap (_cbf );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_abc .AddBitmap (_gfg );_dde ,_cbg :=_gfg .Width ,_gfg .Height ;_ac :=uint64 (_cbg )*uint64 (_dde );_bef .TemplatesSize .Add (_ac ,_gbg );if _cdd ,_gdb =_gbdd .Get (_cbf );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_abc .AddBox (_cdd );_bef .ClassInstances .AddBitmaps (_abc );_bef .CentroidPointsTemplates .AddPoint (_dba ,_eg );_bef .FgTemplates .AddInt (_aad );_bef .UndilatedTemplates .AddBitmap (_gfg );_eaf =(_gbc .Width -2*JbAddedPixels )*(_gbc .Height -2*JbAddedPixels );if _gdb =_bef .TemplateAreas .Add (_eaf );_gdb !=nil {return _c .Wrap (_gdb ,_fba ,"\u0021\u0066\u006f\u0075\u006e\u0064");};};};_bef .NumberOfClasses =len (_bef .UndilatedTemplates .Values );return nil ;};const JbAddedPixels =6;func (_ffe *Classer )classifyRankHaus (_gac *_fc .Boxes ,_fbd *_fc .Bitmaps ,_edfg int )error {const _abg ="\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073";if _gac ==nil {return _c .Error (_abg ,"\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");};if _fbd ==nil {return _c .Error (_abg ,"\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");};_fcb :=len (_fbd .Values );if _fcb ==0{return _c .Error (_abg ,"e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073");};_eecg :=_fbd .CountPixels ();_ggb :=_ffe .Settings .SizeHaus ;_efg :=_fc .SelCreateBrick (_ggb ,_ggb ,_ggb /2,_ggb /2,_fc .SelHit );_ebe :=&_fc .Bitmaps {Values :make ([]*_fc .Bitmap ,_fcb )};_fea :=&_fc .Bitmaps {Values :make ([]*_fc .Bitmap ,_fcb )};var (_fcg ,_dae ,_cfa *_fc .Bitmap ;_ebg error ;);for _acg :=0;_acg < _fcb ;_acg ++{_fcg ,_ebg =_fbd .GetBitmap (_acg );if _ebg !=nil {return _c .Wrap (_ebg ,_abg ,"");};_dae ,_ebg =_fcg .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);if _ebg !=nil {return _c .Wrap (_ebg ,_abg ,"");};_cfa ,_ebg =_fc .Dilate (nil ,_dae ,_efg );if _ebg !=nil {return _c .Wrap (_ebg ,_abg ,"");};_ebe .Values [_fcb ]=_dae ;_fea .Values [_fcb ]=_cfa ;};_abcb ,_ebg :=_fc .Centroids (_ebe .Values );if _ebg !=nil {return _c .Wrap (_ebg ,_abg ,"");};if _ebg =_abcb .Add (_ffe .CentroidPoints );_ebg !=nil {_g .Log .Trace ("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064");};if _ffe .Settings .RankHaus ==1.0{_ebg =_ffe .classifyRankHouseOne (_gac ,_fbd ,_ebe ,_fea ,_abcb ,_edfg );}else {_ebg =_ffe .classifyRankHouseNonOne (_gac ,_fbd ,_ebe ,_fea ,_abcb ,_eecg ,_edfg );};if _ebg !=nil {return _c .Wrap (_ebg ,_abg ,"");};return nil ;};func (_fded *Classer )classifyRankHouseOne (_dfb *_fc .Boxes ,_efe ,_dcc ,_cfde *_fc .Bitmaps ,_ffad *_fc .Points ,_dag int )(_cab error ){const _dgd ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";var (_eeb ,_eda ,_aadf ,_ba float32 ;_gccd int ;_agd ,_eag ,_bag ,_de ,_fdef *_fc .Bitmap ;_bgb ,_gab bool ;);for _edg :=0;_edg < len (_efe .Values );_edg ++{_eag =_dcc .Values [_edg ];_bag =_cfde .Values [_edg ];_eeb ,_eda ,_cab =_ffad .GetGeometry (_edg );if _cab !=nil {return _c .Wrapf (_cab ,_dgd ,"\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079");};_gbcb :=len (_fded .UndilatedTemplates .Values );_bgb =false ;_fdc :=_bdfcf (_fded ,_eag );for _gccd =_fdc .Next ();_gccd > -1;{_de ,_cab =_fded .UndilatedTemplates .GetBitmap (_gccd );if _cab !=nil {return _c .Wrap (_cab ,_dgd ,"\u0062\u006d\u0033");};_fdef ,_cab =_fded .DilatedTemplates .GetBitmap (_gccd );if _cab !=nil {return _c .Wrap (_cab ,_dgd ,"\u0062\u006d\u0034");};_aadf ,_ba ,_cab =_fded .CentroidPointsTemplates .GetGeometry (_gccd );if _cab !=nil {return _c .Wrap (_cab ,_dgd ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073");};_gab ,_cab =_fc .HausTest (_eag ,_bag ,_de ,_fdef ,_eeb -_aadf ,_eda -_ba ,MaxDiffWidth ,MaxDiffHeight );if _cab !=nil {return _c .Wrap (_cab ,_dgd ,"");};if _gab {_bgb =true ;if _cab =_fded .ClassIDs .Add (_gccd );_cab !=nil {return _c .Wrap (_cab ,_dgd ,"");};if _cab =_fded .ComponentPageNumbers .Add (_dag );_cab !=nil {return _c .Wrap (_cab ,_dgd ,"");};if _fded .Settings .KeepClassInstances {_bgc ,_cabf :=_fded .ClassInstances .GetBitmaps (_gccd );if _cabf !=nil {return _c .Wrap (_cabf ,_dgd ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_agd ,_cabf =_efe .GetBitmap (_edg );if _cabf !=nil {return _c .Wrap (_cabf ,_dgd ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_bgc .AddBitmap (_agd );_bdb ,_cabf :=_dfb .Get (_edg );if _cabf !=nil {return _c .Wrap (_cabf ,_dgd ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_bgc .AddBox (_bdb );};break ;};};if !_bgb {if _cab =_fded .ClassIDs .Add (_gbcb );_cab !=nil {return _c .Wrap (_cab ,_dgd ,"");};if _cab =_fded .ComponentPageNumbers .Add (_dag );_cab !=nil {return _c .Wrap (_cab ,_dgd ,"");};_fed :=&_fc .Bitmaps {};_agd ,_cab =_efe .GetBitmap (_edg );if _cab !=nil {return _c .Wrap (_cab ,_dgd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_fed .Values =append (_fed .Values ,_agd );_bbc ,_ecdg :=_agd .Width ,_agd .Height ;_fded .TemplatesSize .Add (uint64 (_ecdg )*uint64 (_bbc ),_gbcb );_gadf ,_beba :=_dfb .Get (_edg );if _beba !=nil {return _c .Wrap (_beba ,_dgd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_fed .AddBox (_gadf );_fded .ClassInstances .AddBitmaps (_fed );_fded .CentroidPointsTemplates .AddPoint (_eeb ,_eda );_fded .UndilatedTemplates .AddBitmap (_eag );_fded .DilatedTemplates .AddBitmap (_bag );};};return nil ;};func (_gbdb *similarTemplatesFinder )Next ()int {var (_cff ,_fef ,_cbfg ,_acc int ;_fad bool ;_dbaa *_fc .Bitmap ;_fbf error ;);for {if _gbdb .Index >=25{return -1;};_fef =_gbdb .Width +TwoByTwoWalk [2*_gbdb .Index ];_cff =_gbdb .Height +TwoByTwoWalk [2*_gbdb .Index +1];if _cff < 1||_fef < 1{_gbdb .Index ++;continue ;};if len (_gbdb .CurrentNumbers )==0{_gbdb .CurrentNumbers ,_fad =_gbdb .Classer .TemplatesSize .GetSlice (uint64 (_fef )*uint64 (_cff ));if !_fad {_gbdb .Index ++;continue ;};_gbdb .N =0;};_cbfg =len (_gbdb .CurrentNumbers );for ;_gbdb .N < _cbfg ;_gbdb .N ++{_acc =_gbdb .CurrentNumbers [_gbdb .N ];_dbaa ,_fbf =_gbdb .Classer .DilatedTemplates .GetBitmap (_acc );if _fbf !=nil {_g .Log .Debug ("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020");return 0;};if _dbaa .Width -2*JbAddedPixels ==_fef &&_dbaa .Height -2*JbAddedPixels ==_cff {return _acc ;};};_gbdb .Index ++;_gbdb .CurrentNumbers =nil ;};};type Classer struct{BaseIndex int ;Settings Settings ;ComponentsNumber *_a .IntSlice ;TemplateAreas *_a .IntSlice ;Widths map[int ]int ;Heights map[int ]int ;NumberOfClasses int ;ClassInstances *_fc .BitmapsArray ;UndilatedTemplates *_fc .Bitmaps ;DilatedTemplates *_fc .Bitmaps ;TemplatesSize _a .IntsMap ;FgTemplates *_a .NumSlice ;CentroidPoints *_fc .Points ;CentroidPointsTemplates *_fc .Points ;ClassIDs *_a .IntSlice ;ComponentPageNumbers *_a .IntSlice ;PtaUL *_fc .Points ;PtaLL *_fc .Points ;};func (_fb *Classer )ComputeLLCorners ()(_ce error ){const _d ="\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _fb .PtaUL ==nil {return _c .Error (_d ,"\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064");};_bb :=len (*_fb .PtaUL );_fb .PtaLL =&_fc .Points {};var (_ef ,_aee float32 ;_fg ,_cf int ;_aa *_fc .Bitmap ;);for _cg :=0;_cg < _bb ;_cg ++{_ef ,_aee ,_ce =_fb .PtaUL .GetGeometry (_cg );if _ce !=nil {_g .Log .Debug ("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_ce );return _c .Wrap (_ce ,_d ,"\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079");};_fg ,_ce =_fb .ClassIDs .Get (_cg );if _ce !=nil {_g .Log .Debug ("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076",_ce );return _c .Wrap (_ce ,_d ,"\u0043l\u0061\u0073\u0073\u0049\u0044");};_aa ,_ce =_fb .UndilatedTemplates .GetBitmap (_fg );if _ce !=nil {_g .Log .Debug ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_ce );return _c .Wrap (_ce ,_d ,"\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_cf =_aa .Height ;_fb .PtaLL .AddPoint (_ef ,_aee +float32 (_cf ));};return nil ;};const (MaxConnCompWidth =350;MaxCharCompWidth =350;MaxWordCompWidth =1000;MaxCompHeight =120;);func Init (settings Settings )(*Classer ,error ){const _ff ="\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074";_ec :=&Classer {Settings :settings ,Widths :map[int ]int {},Heights :map[int ]int {},TemplatesSize :_a .IntsMap {},TemplateAreas :&_a .IntSlice {},ComponentPageNumbers :&_a .IntSlice {},ClassIDs :&_a .IntSlice {},ComponentsNumber :&_a .IntSlice {},CentroidPoints :&_fc .Points {},CentroidPointsTemplates :&_fc .Points {},UndilatedTemplates :&_fc .Bitmaps {},DilatedTemplates :&_fc .Bitmaps {},ClassInstances :&_fc .BitmapsArray {},FgTemplates :&_a .NumSlice {}};if _af :=_ec .Settings .Validate ();_af !=nil {return nil ,_c .Wrap (_af ,_ff ,"");};return _ec ,nil ;};var _ab bool ;func (_fcfg *Settings )SetDefault (){if _fcfg .MaxCompWidth ==0{switch _fcfg .Components {case _fc .ComponentConn :_fcfg .MaxCompWidth =MaxConnCompWidth ;case _fc .ComponentCharacters :_fcfg .MaxCompWidth =MaxCharCompWidth ;case _fc .ComponentWords :_fcfg .MaxCompWidth =MaxWordCompWidth ;};};if _fcfg .MaxCompHeight ==0{_fcfg .MaxCompHeight =MaxCompHeight ;};if _fcfg .Thresh ==0.0{_fcfg .Thresh =0.9;};if _fcfg .WeightFactor ==0.0{_fcfg .WeightFactor =0.75;};if _fcfg .RankHaus ==0.0{_fcfg .RankHaus =0.97;};if _fcfg .SizeHaus ==0{_fcfg .SizeHaus =2;};};func (_ecg Settings )Validate ()error {const _gbgd ="\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065";if _ecg .Thresh < 0.4||_ecg .Thresh > 0.98{return _c .Error (_gbgd ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d");};if _ecg .WeightFactor < 0.0||_ecg .WeightFactor > 1.0{return _c .Error (_gbgd ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d");};if _ecg .RankHaus < 0.5||_ecg .RankHaus > 1.0{return _c .Error (_gbgd ,"\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d");};if _ecg .SizeHaus < 1||_ecg .SizeHaus > 10{return _c .Error (_gbgd ,"\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]");};switch _ecg .Components {case _fc .ComponentConn ,_fc .ComponentCharacters ,_fc .ComponentWords :default:return _c .Error (_gbgd ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et");};return nil ;};func (_bbg *Classer )classifyRankHouseNonOne (_agde *_fc .Boxes ,_add ,_faa ,_aaeb *_fc .Bitmaps ,_gde *_fc .Points ,_fcae *_a .NumSlice ,_ccf int )(_cad error ){const _dec ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";var (_bgbb ,_ddb ,_dfd ,_eae float32 ;_bfa ,_bcf ,_afda int ;_agbf ,_eff ,_efc ,_gbga ,_gffa *_fc .Bitmap ;_ddee ,_gce bool ;);_gfgb :=_fc .MakePixelSumTab8 ();for _ceb :=0;_ceb < len (_add .Values );_ceb ++{if _eff ,_cad =_faa .GetBitmap (_ceb );_cad !=nil {return _c .Wrap (_cad ,_dec ,"b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _bfa ,_cad =_fcae .GetInt (_ceb );_cad !=nil {_g .Log .Trace ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_ceb ,_cad );};if _efc ,_cad =_aaeb .GetBitmap (_ceb );_cad !=nil {return _c .Wrap (_cad ,_dec ,"b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _bgbb ,_ddb ,_cad =_gde .GetGeometry (_ceb );_cad !=nil {return _c .Wrapf (_cad ,_dec ,"\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079");};_fcc :=len (_bbg .UndilatedTemplates .Values );_ddee =false ;_cbb :=_bdfcf (_bbg ,_eff );for _afda =_cbb .Next ();_afda > -1;{if _gbga ,_cad =_bbg .UndilatedTemplates .GetBitmap (_afda );_cad !=nil {return _c .Wrap (_cad ,_dec ,"\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");};if _bcf ,_cad =_bbg .FgTemplates .GetInt (_afda );_cad !=nil {_g .Log .Trace ("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_afda ,_cad );};if _gffa ,_cad =_bbg .DilatedTemplates .GetBitmap (_afda );_cad !=nil {return _c .Wrap (_cad ,_dec ,"\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");};if _dfd ,_eae ,_cad =_bbg .CentroidPointsTemplates .GetGeometry (_afda );_cad !=nil {return _c .Wrap (_cad ,_dec ,"\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d");};_gce ,_cad =_fc .RankHausTest (_eff ,_efc ,_gbga ,_gffa ,_bgbb -_dfd ,_ddb -_eae ,MaxDiffWidth ,MaxDiffHeight ,_bfa ,_bcf ,float32 (_bbg .Settings .RankHaus ),_gfgb );if _cad !=nil {return _c .Wrap (_cad ,_dec ,"");};if _gce {_ddee =true ;if _cad =_bbg .ClassIDs .Add (_afda );_cad !=nil {return _c .Wrap (_cad ,_dec ,"");};if _cad =_bbg .ComponentPageNumbers .Add (_ccf );_cad !=nil {return _c .Wrap (_cad ,_dec ,"");};if _bbg .Settings .KeepClassInstances {_dcf ,_gfd :=_bbg .ClassInstances .GetBitmaps (_afda );if _gfd !=nil {return _c .Wrap (_gfd ,_dec ,"\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029");};if _agbf ,_gfd =_add .GetBitmap (_ceb );_gfd !=nil {return _c .Wrap (_gfd ,_dec ,"\u0070i\u0078\u0061\u005b\u0069\u005d");};_dcf .Values =append (_dcf .Values ,_agbf );_ffbe ,_gfd :=_agde .Get (_ceb );if _gfd !=nil {return _c .Wrap (_gfd ,_dec ,"b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};_dcf .Boxes =append (_dcf .Boxes ,_ffbe );};break ;};};if !_ddee {if _cad =_bbg .ClassIDs .Add (_fcc );_cad !=nil {return _c .Wrap (_cad ,_dec ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _cad =_bbg .ComponentPageNumbers .Add (_ccf );_cad !=nil {return _c .Wrap (_cad ,_dec ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_baa :=&_fc .Bitmaps {};_agbf =_add .Values [_ceb ];_baa .AddBitmap (_agbf );_bfc ,_dbb :=_agbf .Width ,_agbf .Height ;_bbg .TemplatesSize .Add (uint64 (_bfc )*uint64 (_dbb ),_fcc );_cde ,_ddc :=_agde .Get (_ceb );if _ddc !=nil {return _c .Wrap (_ddc ,_dec ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_baa .AddBox (_cde );_bbg .ClassInstances .AddBitmaps (_baa );_bbg .CentroidPointsTemplates .AddPoint (_bgbb ,_ddb );_bbg .UndilatedTemplates .AddBitmap (_eff );_bbg .DilatedTemplates .AddBitmap (_efc );_bbg .FgTemplates .AddInt (_bfa );};};_bbg .NumberOfClasses =len (_bbg .UndilatedTemplates .Values );return nil ;};