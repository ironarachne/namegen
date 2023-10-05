package namegen

var (
	anglosaxonMaleFirstNames = []string{
		"Ælle",
		"Ælfwine",
		"Æðelbert",
		"Æðelfrid",
		"Æðelhun",
		"Æðelred",
		"Æðelwald",
		"Æðelwalh",
		"Alchfrid",
		"Aldfrid",
		"Aldhelm",
		"Alduini",
		"Aldwulf",
		"Alric",
		"Andhun",
		"Anna",
		"Berðun",
		"Bertwald",
		"Cædmon",
		"Cælin",
		"Cearl",
		"Ceolwulf",
		"Coenred",
		"Coenwalh",
		"Cuichelm",
		"Cuðbert",
		"Cynegels",
		"Cynefrid",
		"Eadbald",
		"Eadbert",
		"Eadfrid",
		"Eadric",
		"Eadwyn",
		"Eafa",
		"Eanfrid",
		"Earconbert",
		"Earpwald",
		"Egbert",
		"Egfrid",
		"Egric",
		"Eni",
		"Hengist",
		"Hereric",
		"Hlothere",
		"Horsa",
		"Hunwald",
		"Immin",
		"Irminric",
		"Octa",
		"Oeric",
		"Offa",
		"Osfrid",
		"Osred",
		"Osric",
		"Oswald",
		"Oswy",
		"Oswyn",
		"Peada",
		"Penda",
		"Rægenhere",
		"Redwald",
		"Ricbert",
		"Sabert",
		"Sebbi",
		"Sigbert",
		"Sighard",
		"Sighere",
		"Swefred",
		"Swidhelm",
		"Thrydwulf",
		"Tondbert",
		"Tondhere",
		"Wictred",
		"Wilfrid",
		"Wini",
		"Wuffa",
		"Wuscfrea",
		"Wulfhere",
		"Yffi",
	}

	anglosaxonFemaleFirstNames = []string{
		"Acha",
		"Ælffled",
		"Æðelberga",
		"Æðeldreda",
		"Æðelðryd",
		"Æðelhild",
		"Breguswið",
		"Coenberg",
		"Cyneburg",
		"Cynwise",
		"Eabæ",
		"Eanfled",
		"Earcongota",
		"Edyð",
		"Hilda",
		"Hereswið",
		"Osðryd",
		"Tortgyð",
		"Sexburg",
	}

	anglosaxonLastNames = []string{
		"Ackart",
		"Ackerman",
		"Ackers",
		"Ackland",
		"Ackman",
		"Acton",
		"Ada",
		"Adlam",
		"Adolphus",
		"Aiken",
		"Aikman",
		"Akeman",
		"Aken",
		"Akers",
		"Akin",
		"Albrecht",
		"Alden",
		"Aldersey",
		"Aldis",
		"Aldjoy",
		"Aldred",
		"Aldridge",
		"Alford",
		"Alfred",
		"Alice",
		"Allgood",
		"Alvin",
		"Alvord",
		"Alwin",
		"Amherst",
		"Armistead",
		"Armsted",
		"Ashby",
		"Ashford",
		"Ashley",
		"Ashton",
		"Askew",
		"Aspinwall",
		"Astley",
		"Aston",
		"Atherton",
		"Audley",
		"Bagley",
		"Baker",
		"Barclay",
		"Barnard",
		"Barras",
		"Barton",
		"Bateman",
		"Bath",
		"Bathurst",
		"Beckford",
		"Bedford",
		"Beorn",
		"Berkeley",
		"Bernard",
		"Bertram",
		"Bertrand",
		"Biddulph",
		"Bogue",
		"Boorman",
		"Bowers",
		"Bristed",
		"Bristow",
		"Brock",
		"Buckingham",
		"Buckminster",
		"Buckston",
		"Burby",
		"Burden",
		"Burleigh",
		"Burr",
		"Buxton",
		"Byington",
		"Denton",
		"Dering",
		"Mansfield",
		"Mildmay",
		"Milton",
		"Minster",
		"Mitchell",
		"Moseley",
		"Moxley",
		"Musgrave",
		"Newbury",
		"Nott",
		"Ockley",
		"Ogden",
		"Osborn",
		"Osmund",
		"Oswald",
		"Canning",
		"Chalk",
		"Channing",
		"Chatham",
		"Chatsey",
		"Chatterton",
		"Dillingham",
		"Dinton",
		"Dixie",
		"Ran",
		"Randal",
		"Randolph",
		"Randulph",
		"Reed",
		"Reynolds",
		"Reynoldson",
		"Richmond",
		"Roberts",
		"Rodland",
		"Roland",
		"Rollin",
		"Rowley",
		"Salisbury",
		"Sarisbury",
		"Saterlee",
		"Scarborough",
		"Scardsdale",
		"Sebright",
		"Sherwood",
		"Skeffington",
		"Skelton",
		"Stanwood",
		"Chetham",
		"Chilton",
		"Chubb",
		"Dunning",
		"Dunstan",
		"Durham",
		"Eaton",
		"Eddy",
		"Edgar",
		"Ediker",
		"Edmond",
		"Edward",
		"Edwards",
		"Egbert",
		"Eldred",
		"Emerson",
		"Ethelbert",
		"Everard",
		"Eytinge",
		"Fagg",
		"Fairfax",
		"Falkland",
		"Stocking",
		"Theobald",
		"Walden",
		"Waldgrave",
		"Waldron",
		"Wales",
		"Wallace",
		"Waller",
		"Whiting",
		"Whitlock",
		"Whitney",
		"Wickham",
		"Wickliff",
		"Wilder",
		"Winchcombe",
		"Winship",
		"Coop",
		"Coope",
		"Coops",
		"Cope",
		"Copp",
		"Coppen",
		"Francis",
		"Frankland",
		"Giffard",
		"Gifford",
		"Goodrich",
		"Goss",
		"Granger",
		"Halden",
		"Halifax",
		"Halsey",
		"Hanna",
		"Hansel",
		"Harcourt",
		"Hargrave",
		"Harleigh",
		"Harley",
		"Harlow",
		"Harvey",
		"Haw",
		"Hawes",
		"Hawley",
		"Hayman",
		"Hayward",
		"Heaton",
		"Henry",
		"Herne",
		"Hewer",
		"Heyman",
		"Hildyard",
		"Hilyard",
		"Hippisley",
		"Holcombe",
		"Holtcombe",
		"Worth",
		"Yare",
		"Yule",
		"Cutting",
		"Hood",
		"Hopper",
		"Hotham",
		"Houghton",
		"Hubert",
		"Huntington",
		"Hurst",
		"Husted",
		"Kemp",
		"Kendrick",
		"Kenward",
		"Lambert",
		"Lawley",
		"Lind",
		"Litchfield",
	}
)