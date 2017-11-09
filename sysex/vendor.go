// generate: stringer -type VendorID
package sysex

type VendorID uint

func (v VendorID) Bytes() []byte {
	if v&0xFF0000 != 0 {
		return []byte{byte((v >> 16) & 0xFF)}
	}
	return []byte{
		0x00,
		byte((v >> 8) & 0xFF),
		byte(v & 0xFF),
	}
}

func vendorFrom3Bytes(b []byte) VendorID {
	ui0 := uint(b[0])
	ui1 := uint(b[1])
	ui2 := uint(b[2])

	return VendorID(ui0<<16 | ui1<<8 | ui2)
}

func vendorFrom1Byte(b []byte) VendorID {
	return VendorID((uint(b[0]) << 16) & 0xFF0000)
}

const (
	NonCommercialSystemExclusive      VendorID = 0x7D0000
	UniversalSystemExclusive          VendorID = 0x7E0000
	RealTimeUniversalSystemExclusive  VendorID = 0x7F0000
	Sequential                        VendorID = 0x010000
	IDP                               VendorID = 0x020000
	VoyetraOctavePlateau              VendorID = 0x030000
	Moog                              VendorID = 0x040000
	PassportDesigns                   VendorID = 0x050000
	Lexicon                           VendorID = 0x060000
	Kurzweil                          VendorID = 0x070000
	Fender                            VendorID = 0x080000
	Gulbransen                        VendorID = 0x090000
	AKGAcoustics                      VendorID = 0x0A0000
	VoyceMusic                        VendorID = 0x0B0000
	WaveframeCorp                     VendorID = 0x0C0000
	ADASignalProcessors               VendorID = 0x0D0000
	GarfieldElectronics               VendorID = 0x0E0000
	Ensoniq                           VendorID = 0x0F0000
	Oberheim                          VendorID = 0x100000
	AppleComputer                     VendorID = 0x110000
	GreyMatterResponse                VendorID = 0x120000
	Digidesign                        VendorID = 0x130000
	PalmTreeInstruments               VendorID = 0x140000
	JLCooperElectronics               VendorID = 0x150000
	Lowrey                            VendorID = 0x160000
	AdamsSmith                        VendorID = 0x170000
	EmuSystems                        VendorID = 0x180000
	HarmonySystems                    VendorID = 0x190000
	ART                               VendorID = 0x1A0000
	Baldwin                           VendorID = 0x1B0000
	Eventide                          VendorID = 0x1C0000
	Inventronics                      VendorID = 0x1D0000
	Clarity                           VendorID = 0x1F0000
	Passac                            VendorID = 0x200000
	SIEL                              VendorID = 0x210000
	Synthaxe                          VendorID = 0x220000
	Hohner                            VendorID = 0x240000
	Twister                           VendorID = 0x250000
	Solton                            VendorID = 0x260000
	JellinghausMS                     VendorID = 0x270000
	SouthworthMusicSystems            VendorID = 0x270000
	PPG                               VendorID = 0x290000
	JEN                               VendorID = 0x2A0000
	SSLLimited                        VendorID = 0x2B0000
	AudioVeritrieb                    VendorID = 0x2C0000
	Elka                              VendorID = 0x2F0000
	Dynacord                          VendorID = 0x300000
	Viscount                          VendorID = 0x310000
	ClaviaDigitalInstruments          VendorID = 0x330000
	AudioArchitecture                 VendorID = 0x340000
	GeneralMusicCorp                  VendorID = 0x350000
	SoundcraftElectronics             VendorID = 0x390000
	Wersi                             VendorID = 0x3B0000
	AvabElectronikAb                  VendorID = 0x3C0000
	Digigram                          VendorID = 0x3D0000
	WaldorfElectronics                VendorID = 0x3E0000
	Quasimidi                         VendorID = 0x3F0000
	Kawai                             VendorID = 0x400000
	Roland                            VendorID = 0x410000
	Korg                              VendorID = 0x420000
	Yamaha                            VendorID = 0x430000
	Casio                             VendorID = 0x440000
	KamiyaStudio                      VendorID = 0x460000
	Akai                              VendorID = 0x470000
	JapanVictor                       VendorID = 0x480000
	Mesosha                           VendorID = 0x490000
	HoshinoGakki                      VendorID = 0x4A0000
	FujitsuElect                      VendorID = 0x4B0000
	Sony                              VendorID = 0x4C0000
	NisshinOnpa                       VendorID = 0x4D0000
	TEAC                              VendorID = 0x4E0000
	MatsushitaElectric                VendorID = 0x500000
	Fostex                            VendorID = 0x510000
	Zoom                              VendorID = 0x520000
	MidoriElectronics                 VendorID = 0x530000
	MatsushitaCommunicationIndustrial VendorID = 0x540000
	SuzukiMusicalInstMfg              VendorID = 0x550000
	TimeWarnerInteractive             VendorID = 0x000001
	DigitalMusicCorp                  VendorID = 0x000007
	IOTASystems                       VendorID = 0x000008
	NewEnglandDigital                 VendorID = 0x000009
	Artisyn                           VendorID = 0x00000A
	IVLTechnologies                   VendorID = 0x00000B
	SouthernMusicSystems              VendorID = 0x00000C
	LakeButlerSoundCompany            VendorID = 0x00000D
	Alesis                            VendorID = 0x00000E
	DODElectronics                    VendorID = 0x000010
	StuderEditech                     VendorID = 0x000011
	PerfectFretworks                  VendorID = 0x000014
	KAT                               VendorID = 0x000015
	Opcode                            VendorID = 0x000016
	RaneCorp                          VendorID = 0x000017
	AnadiInc                          VendorID = 0x000018
	KMX                               VendorID = 0x000019
	AllenAndHeathBrenell              VendorID = 0x00001A
	PeaveyElectronics                 VendorID = 0x00001B
	// 360 Systems
	ThreeSixtySystems            VendorID = 0x00001C
	SpectrumDesignandDevelopment VendorID = 0x00001D
	MarquisMusic                 VendorID = 0x00001E
	ZetaSystems                  VendorID = 0x00001F
	Axxes                        VendorID = 0x000020
	Orban                        VendorID = 0x000021
	KTI                          VendorID = 0x000024
	BreakawayTechnologies        VendorID = 0x000025
	CAE                          VendorID = 0x000026
	RocktronCorp                 VendorID = 0x000029
	PianoDisc                    VendorID = 0x00002A
	CannonResearchGroup          VendorID = 0x00002B
	RogersInstrumentCorp         VendorID = 0x00002D
	BlueSkyLogic                 VendorID = 0x00002E
	EncoreElectronics            VendorID = 0x00002F
	Uptown                       VendorID = 0x000030
	Voce                         VendorID = 0x000031
	// CTI Audio Inc (Music Intel Dev.)
	CTIAudioInc                 VendorID = 0x000032
	SAndSResearch               VendorID = 0x000033
	BroderbundSoftwareInc       VendorID = 0x000034
	AllenOrganCo                VendorID = 0x000035
	MusicQuest                  VendorID = 0x000037
	APHEX                       VendorID = 0x000038
	GallienKrueger              VendorID = 0x000039
	IBM                         VendorID = 0x00003A
	HotzInstrumentsTechnologies VendorID = 0x00003C
	ETALighting                 VendorID = 0x00003D
	NSICorporation              VendorID = 0x00003E
	AdLibInc                    VendorID = 0x00003F
	RichmondSoundDesign         VendorID = 0x000040
	Microsoft                   VendorID = 0x000041
	TheSoftwareToolworks        VendorID = 0x000042
	NicheRJMG                   VendorID = 0x000043
	Intone                      VendorID = 0x000044
	GTElectronicsGrooveTubes    VendorID = 0x000047
	TimelineVista               VendorID = 0x000049
	MesaBoogie                  VendorID = 0x00004A
	SequoiaDevelopment          VendorID = 0x00004C
	StudioElectronics           VendorID = 0x00004D
	Euphonix                    VendorID = 0x00004E
	InterMIDI                   VendorID = 0x00004F
	MIDISolutions               VendorID = 0x000050
	// 3DO Company
	ThreeDOCompany     VendorID = 0x000051
	LightwaveResearch  VendorID = 0x000052
	MicroW             VendorID = 0x000053
	SpectralSynthesis  VendorID = 0x000054
	LoneWolf           VendorID = 0x000055
	StudioTechnologies VendorID = 0x000056
	PetersonEMP        VendorID = 0x000057
	Atari              VendorID = 0x000058
	MarionSystems      VendorID = 0x000059
	DesignEvent        VendorID = 0x00005A
	WinjammerSoftware  VendorID = 0x00005B
	ATTBellLabs        VendorID = 0x00005C
	// AT&T Bell Labs
	Symetrix                VendorID = 0x00005E
	MIDItheWorld            VendorID = 0x00005F
	DesperProducts          VendorID = 0x000060
	MicrosNMIDI             VendorID = 0x000061
	AccordiansIntl          VendorID = 0x000062
	EuPhonics               VendorID = 0x000063
	Musonix                 VendorID = 0x000064
	TurtleBeachSystems      VendorID = 0x000065
	MackieDesigns           VendorID = 0x000066
	Compuserve              VendorID = 0x000067
	BESTechnologies         VendorID = 0x000068
	QRSMusicRolls           VendorID = 0x000069
	PGMusic                 VendorID = 0x00006A
	SierraSemiconductor     VendorID = 0x00006B
	EpiGrafAudioVisual      VendorID = 0x00006C
	ElectronicsDeiversified VendorID = 0x00006D
	Tune1000                VendorID = 0x00006E
	AdvancedMicroDevices    VendorID = 0x00006F
	Mediamation             VendorID = 0x000070
	SabineMusic             VendorID = 0x000071
	WoogLabs                VendorID = 0x000072
	Micropolis              VendorID = 0x000073
	TaHorngMusicalInst      VendorID = 0x000074
	// eTek was formerly known as forte
	ETekForte                 VendorID = 0x000075
	Electrovoice              VendorID = 0x000076
	Midisoft                  VendorID = 0x000077
	QSoundLabs                VendorID = 0x000078
	Westrex                   VendorID = 0x000079
	NVidia                    VendorID = 0x00007A
	ESSTechnology             VendorID = 0x00007B
	MediaTrixPeripherals      VendorID = 0x00007C
	Brooktree                 VendorID = 0x00007D
	Otari                     VendorID = 0x00007E
	KeyElectronics            VendorID = 0x00007F
	CrystalakeMultimedia      VendorID = 0x000080
	CrystalSemiconductor      VendorID = 0x000081
	RockwellSemiconductor     VendorID = 0x000081
	Dream                     VendorID = 0x002000
	StrandLighting            VendorID = 0x002001
	AmekSystems               VendorID = 0x002002
	BohmElectronic            VendorID = 0x002004
	TridentAudio              VendorID = 0x002006
	RealWorldStudio           VendorID = 0x002007
	YesTechnology             VendorID = 0x002009
	Audiomatica               VendorID = 0x00200A
	BontempiFarfisa           VendorID = 0x00200B
	FBTElettronica            VendorID = 0x00200C
	MidiTemp                  VendorID = 0x00200D
	LAAudio                   VendorID = 0x00200E
	Zero88LightingLimited     VendorID = 0x00200F
	MiconAudioElectronicsGmbH VendorID = 0x002010
	ForefrontTechnology       VendorID = 0x002011
	KentonElectronics         VendorID = 0x002013
	ADB                       VendorID = 0x002015
	MarshallProducts          VendorID = 0x002016
	DDA                       VendorID = 0x002017
	BSS                       VendorID = 0x002018
	MALightingTechnology      VendorID = 0x002018
	Fatar                     VendorID = 0x00201A
	QSCAudio                  VendorID = 0x00201B
	ArtisanClassicOrgan       VendorID = 0x00201C
	OrlaSpa                   VendorID = 0x00201D
	PinnacleAudio             VendorID = 0x00201E
	TCElectonics              VendorID = 0x00201F
	DoepferMusikelektronik    VendorID = 0x002020
	CreativeTechnologyPte     VendorID = 0x002021
	MinamiSeiyddo             VendorID = 0x002022
	Goldstar                  VendorID = 0x002023
	MidisoftsasdiMCima        VendorID = 0x002024
	Samick                    VendorID = 0x002025
	PennyandGiles             VendorID = 0x002026
	AcornComputer             VendorID = 0x002027
	LSCElectronics            VendorID = 0x002028
	NovationEMS               VendorID = 0x002029
	SamkyungMechatronics      VendorID = 0x00202A
	MedeliElectronics         VendorID = 0x00202B
	CharlieLab                VendorID = 0x00202C
	BlueChipMusicTech         VendorID = 0x00202D
	BEEOHCorp                 VendorID = 0x00202E
)
