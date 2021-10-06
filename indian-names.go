package namegen

var (
	indianMaleFirstNames = []string{
		"Aadarsh", "Aadesh", "Aadhidev", "Aadhishankar", "Aadit","Aaditey", "Aagman", "Aagney", "Aahva",
		"Aahwaanith", "Aakarshan", "Aamod","Aarnav", "Abhay","Abhi","Abhijat", "Abhijay", "Abhijit /Abhijeet", "Abhik",
		"Abhilash", "Abhinandan", "Abhinav", "Abhindra", "Abhineeth", "Abhinivesh", "Abhiraj", "Abhiram", "Abhiroop", 
		"Abhirut", "Abhisar", "Abhishek", "Abhyas", "Abra","Achal",
		"Achalraj", "Achintya", "Achlendra", "Achyuta", "Adamya", "Adarsh", "Adeel","Adeep",
		"Adheesh", "Adhip","Adi","Adinath", "Adish","Adith","Aditya", "Aditya vardhan", "Adjovi", "Adrut","Advait", "Advay","Advik",
		"Agastya", "Agendra", "Aghosh", "Agneya", "Agnivesh", "Agraj","Agrim","Ahan","Ainesh", "Ajar","Ajatashatru", "Ajatashatru", "Ajay","Ajit",
		"Ajitabh", "Ajkhyat", "Ajmil","Akaash", "Akarsh", "Akav","Akhilesh", "Aksat",
		"Akshan", "Akshat", "Akshat", "Akshay", "Akshayakeerti", "Akshit", "Akshobhya", "Akshya", "Akul",
		"Akyhielan", "Alaap","Alagan", "Alankar", "Alok","Alok Nath", "Alop","Amaan","Amal","Amalendu", "Amalesh", "Amanuday", "Amar","Amara",
		"Amardeep", "Amarjeet", "Amav","Amay","Ambar","Ambarisha", "Ambikapathi", "Ambud","Ambuj","Ameet",
		"Ameya","Amin","Amirdeswaran", "Amish","Amit","Amitaabh", "Amitabh", "Amitava", "Amitesh", "Amitosh", "Amod","Amogh",
		"Amol","Amraj","Amrik","Amrish", "Amul","Amulaya", "Amulya", "Anagh","Anan","Anand","Anandmoorti", "Anang","Anant",
		"Anantram", "Ananyo", "Anay","Aneesh", "Angad","Angada", "Angak","Anik","Aniketh", "Anil","Anilkumar", "Animesh", "Aniruddh", "Anirudhh", "Anirudhha", "Anish","Anish",
		"Anit","Anjal","Anjalika", "Anjan","Anjuman", "Ankit","Ankur","Ankush", "Anniruddha", "Anoop / Anup", "Anshu","Anshul", "Anshuman", "Ansu",
		"Antariksh", "Antriksa", "Anubhav", "Anuj","Anupam", "Anurag", "Anvay","Apoorva / Apurva", "Aradhak", "Aradhya", "Arav","Archa",
		"Archit", "Arghya", "Arham","Arihan", "Arihant", "Arin","Ariv","Arjun","Arman","Arnav","Arni",
		"Arulsyankar", "Arun","Arvind", "Aryan","Aryanathan", "Ashir","Ashirvad", "Ashish", "Ashish blessings", "Ashmit", "Ashok",
		"Ashray", "Ashumu", "Ashutosh", "Ashvath", "Ashvin", "Ashwath", "Ashwathama", "Ashwatthama", "Ashwin", "Asija","Askha",
		"Asuman", "Asvathama", "Aswad","Atal","Ateet","Atharvan", "Athervan", "Athithya", "Athulya", "Atiksh", "Atin","Atmajyoti", "Atman",
		"Atul","Atulya", "Avalok", "Avanindra", "Avanish", "Avichal", "Avinash", "Aviraj", "Aviral", "Avish",
		"Avkash", "Avneesh", "Avnendra", "Ayush","Babala", "Babul","Badal","Badri prasad", "Badrinath", "Baiju",
		"Balaaditya", "Balachandra", "Balaji", "Balakrishna", "Balamani", "Balamohan", "Balashankar", "Balavan", "Baldev", "Balendu", "Bali",
		"Balraj", "Balram", "Balveer", "Balwan", "Baneet", "Bankim", "Bankimchandra", "Bansi",
		"Bansilal", "Barindra", "Basant", "Basdev", "Bhadrak", "Bhadrakapil", "Bhadraksh", "Bhadrashree", "Bhadresh", "Bhagat", "Bhagesh", "Bhagirath", "Bhagwant", "Bhairav", "Bhamhaghosh", "Bhamhanand", "Bhanu",
		"Bhanudas", "Bhanumitra", "Bhanuprakash", "Bhanuprasad", "Bharadwaj", "Bharat", "Bhargav", "Bhartesh", "Bhaskar", "Bhavesh", "Bhavik", "Bhawanidas", "Bhim",
		"Bhishma", "Bhisma", "Bholanath", "Bhoumik", "Bhrigu", "Bhudev", "Bhumin", "Bhumindra", "Bhupal", "Bhupati", "Bhupen", "Bhupendra", "Bhupesh", "Bhushan", "Bhuvan", "Bhuvan", "Bhuvanesh", "Bhuvanesh", "Bijoyanada", "Bikram", "Bindusar", "Byju",
		"Chaitanya", "Chakradhar", "Chakrapani", "Chakshan", "Chaman", "Champak", "Chandan", "Chandarmouli", "Chander moon", "Chandrabhushan", "Chandradatt", "Chandradhar", "Chandragupt", "Chandrahaas", "Chandrahas", "Chandraj", "Chandrak", "Chandrakant", "Chandrakiran", "Chandramouli", "Chandranath", "Chandranshu", "Chandraprakash", "Chandrasen", "Chandrashekhar", "Chandresh", "Chapal", "Charan", "Charandas", "Charanjit", "Charudutta", "Chaturanan", "Chaturvedi", "Chetan", "Chetas", "Chhailbehari", "Chidambar", "Chidambaram", "Chidanand", "Chidhatma", "Chinar", "Chinmay", "Chintamani", "Chintan", "Chintav", "Chintu", "Chirag", "Chiranjeet, Chirayu", "Chiranjeev", "Chirayu", "Chitragupta", "Chitraksh", "Chittaranjan", "Chittaswarup", "Chittesh", "Chrirag", "Chudamani", "Chunmay", "Dahana", "Daksh",
		"Dakshesh", "Dakshin", "Daljeet", "Daman",
		"Damodar", "Daniel", "Darshan", "Darshit", "Daruka", "Dasharath", "Datta",
		"Dattatreya", "Daulat", "Daya",
		"Dayanand", "Dayanidhi", "Dayasagar", "Dayaswarup", "Deenbabdhu", "Deendayal", "Deepak", "Deepanath", "Deepankar", "Deepdas", "Deependra", "Deependu", "Deepesh", "Deeptanshu", "Deeptendu", "Dehabhuj", "Devadutt", "Devam",
		"Devanand", "Devang", "Devansh", "Devaraj", "Devarsh", "Devarshi", "Devashish", "Devdas", "Devendra", "Devesh", "Devguru", "Devidas", "Devilal", "Devin",
		"Deviprasad", "Devitri", "Devkinandan", "Devraj", "Devrity", "Deyaan", "Dhananad", "Dhananjay", "Dhananjay Arjuna", "Dhanaraj", "Dhanesh", "Dhanu",
		"Dhanvantari", "Dharesh", "Dharma", "Dharma", "Dharmadev", "Dharmakeerti", "Dharmavira", "Dharmendra", "Dharmesh", "Dharmik", "Dharuna", "Dhatri", "Dhaval", "Dheeraj", "Dhigana", "Dhiren", "Dhirendra", "Dhiwyannshu", "Dhruv",
		"Dhruva", "Dhumal", "Dhureen", "Dhyanam", "Digamber", "Dikshan", "Dikshil", "Dilanesh", "Diler Brave", "Dilip",
		"Diliso", "Dinanath", "Dinar",
		"Dinesh", "Dinkar", "Dinpal", "Dipen",
		"Dipinder", "Dipten", "Divakar", "Divit",
		"Divjot", "Divyanshu", "Divyendu", "Divyesh", "Druvan", "Duranjaya", "Durga",
		"Durgadas", "Durgadutt", "Durgesh", "Durjaya", "Durmada", "Durvish", "Dvimidha", "Dwarkadhish", "Dwarkanath", "Dwijendra/Dwijesh", "Dyaus",
		"Eashan", "Ekachakra", "Ekalinga", "Ekanga", "Eklavya", "Eknath", "Esh",
		"Eshaan", "Eshwar", "Eshwardutt", "Falak",
		"Fateh",
		"Gaangey", "Gadaadhar", "Gagan",
		"Gaganavihaaree", "Gagandeep", "Gajaanan", "Gajanan", "Gajanand", "Gajbaahu", "Gajendra", "Gajendranath", "Gajkaran", "Gajpati", "Gajvadan", "Gaman",
		"Gambheer", "Ganak",
		"Ganapati", "Ganaraj", "Gandharaj", "Gandharv", "Gandharva", "Gandhik", "Ganesh", "Gangaadhar", "Gangadhar", "Gangadutt", "Gangesh", "Gangol", "Gannaath", "Garg",
		"Garjan", "Garv",
		"Gatik",
		"Gaurang", "Gaurav", "Gaurikant", "Gaurinandan", "Gaurinath", "Gaurish", "Gaurishankar", "Gautam", "Geet",
		"Ghanashyam", "Ghandeep", "Ghanendra", "Gharshit", "Ghoshal", "Giridhar", "Girijanandan", "Girilal", "Giriraj", "Girish", "Gogula", "Gopal",
		"Gopan",
		"Goshanraj", "Gouresh", "Govind", "Gowshik", "Gul",
		"Gulab",
		"Gulal",
		"Gulshan", "Gunaratna", "Gunjal", "Gurbachan", "Gurcharan", "Gurdayal", "Gurdeep", "Gurjas", "Gurkiran", "Gurman", "Gurmeet", "Gurnam", "Gursharan", "Guru",
		"Gurudas", "Gurudutt", "Hansaraj", "Hanshal", "Hansraj", "Hanuman", "Hardik", "Haresh", "Hari",
		"Haridutt", "Harihar", "Harikiran", "Harilal", "Harina", "Harish", "Harishankar", "Harith", "Harman", "Harmendra", "Harsh",
		"Harsh joy", "Harsha", "Harshad", "Harshal", "Harshul", "Harshvardhan", "Harsith", "Harteij", "Hastin", "Havish", "Heet",
		"Hemachandra", "Hemadri", "Hemal",
		"Hemang", "Hemant", "Hemaprakash", "Hemaraj", "Hemavatinandan", "Hemendra", "Hemendu", "Heramba", "Himal",
		"Himank", "Himanshu", "Himmat", "Hiran",
		"Hiranmay", "Hiranya", "Hiresh", "Hiten",
		"Hitendra", "Hitesh", "Hridayesh", "Hridaynath", "Hridyanshu", "Hrishikesh", "Hrithik", "Hrithikesh", "Hrydesh", "Ilesh",
		"Ina",
		"Indeever", "Indiresh", "Indra",
		"Indradutt", "Indrajit", "Indraneel", "Indranil", "Indubhushan", "Induj",
		"Indukant", "Indushekhar", "Inesh",
		"Iravan", "Iresh",
		"Isha",
		"Ishan",
		"Ishana", "Ishir",
		"Ishpreet", "Ishranth", "Ishrit", "Ishver", "Ishwar", "Jagad",
		"Jagadeep", "Jagadeesh", "Jagadev", "Jagadish", "Jagajeet", "Jagajeevan", "Jagamohan", "Jagan",
		"Jaganarayan", "Jagannath", "Jagatbehari", "Jagatkishor", "Jagatpal", "Jagatprabhu", "Jagatprakash", "Jagatprana", "Jagatveer", "Jahi",
		"Jaichand", "Jaideep", "Jaidev", "Jaikirti", "Jaikrishna", "Jainil", "Jaisinha", "Jaisukh", "Jaithra", "Jakarious", "Jalavirya", "Jalbhushan", "Jaldev", "Jaldhar", "Jalendu", "Jamadagni", "Janak",
		"Janakibhushan", "Janakinath", "Janakiraman", "Janam",
		"Janardan", "Janav",
		"Janesh", "Janith", "Janmesh", "Janya",
		"Japan",
		"Japendu", "Japesh", "Jarnu",
		"Jasapal", "Jashan", "Jashith", "Jaskaran", "Jasraj", "Jasveer", "Jaswant", "Jatan",
		"Jatin",
		"Javesh", "Jawahar", "Jay",
		"Jayachand", "Jayadeep", "Jayaditya", "Jayant", "Jayapal", "Jayaprakash", "Jayashekhar", "Jayawant", "Jayesh", "Jaypal", "Jaysukh", "Jeet",
		"Jeethesh", "Jeevan", "Jilesh", "Jimuta", "Jinendra", "Jinesh", "Jishnu", "Jiten",
		"Jitendra", "Jivana", "Jivin",
		"Jnyandeep", "Jnyaneshwar", "Jogindra", "Jograj", "Josh","Jovan",
		"Jwalaprasad", "Jyotiprakash", "Jyotiranjan", "Jyotirdhar", "Jyotirmaya", "Jyotisa", "Jyran",
		"Kailash", "Kailasnath", "Kairav", "Kaish",
		"Kaladhar", "Kalan",
		"Kalanath", "Kalanidhi", "Kalicharan", "Kalidas", "Kalikesh", "Kalkin", "Kalpak", "Kalpanath", "Kalyan", "Kamadev", "Kamal",
		"Kamal Nath", "Kamalakar", "Kamalaksh", "Kamalbandhu", "Kamalekshan", "Kamalesh", "Kamalkant", "Kamalnath", "Kamalnayan", "Kamboj", "Kamlapati", "Kamlesh",
		 "Kamran", "Kanai","Kanaiyya", "Kanak","Kanan","Kanchan", "Kangana", "Kanha",
		"Kanhai", "Kanhaiya", "Kanishk", "Kanth","Kanthamani", "Kapaali", "Kapi","Kapidhwaj", "Kapil","Kapilashwa", "Kapindra", "Kapirath", "Karan",
		"Kardama", "Karna, Karan", "Karnabhushan", "Karnajeet", "Karnapriya", "Kartheek", "Kartik", "Kartik", "Kartikey", "Kartikeya", "Kartikeya Subramanyam", "Karun",
		"Karunakar", "Karunanidhi", "Karunesh", "Kashi Nath", "Kashinath", "Kathit", "Kaushal", "Kaushik", "kausthubh", "Kauthuk", "Kavana", "Kavi",
		"Kavish", "Kayaan", "Kaylor", "Kedar","Kedarnath", "Keertan", "Keerthinath", "Kerak","Kesav","Keshav", "Ketak",
		"Ketan",	"Ketubh", "Keval","Keyan","Khagendra", "Khagesh", "Khairiya", "Kharanshu", "Khemprakash", "Khush","Khushwant", "Kiaan",	"Kiash",
		"Kinnar", "Kinshuk", "Kintan", "Kirit","Kiritmani", "Kirtan", "Kirtibhushan", "Kirtivallabh", "Kishan", "Kishlaya", "Kishore", "Kisna",
		"Koundinya", "Kripa","Kripanidhi", "Kripasagar", "Krishav", "Krishna", "Krishna Das", "Krishnamoorti", "Krithik", "Krunal", "Krupal", 
		"Kshiraj", "Kshithiraj", "Kshitidhar", "Kshitij", "Kuchela", "Kuldeep", "Kuldev", "Kulvir", "Kumudaksh", "Kumudesh", "Kunal",
		"Kundan", "Kunjabehari", "Kunwar", "Kusagra", "Kush",
		"Kushad", "Kushagra", "Kushal", "Kushan", "Kusumakar", "Lailesh", "Lakshit", "Lakshman", "Lakshmigopal", "Lakshmikant", "Lakshminarayan",
		 "Lakshminath", "Lakshmipati", "Lakshmiraman", "Lakshya", "Lalchand,", "Lalit",
		"Lalitaditya", "Lalitchandra", "Lalitesh", "Lalitlochan", "Lalitmohan", "Lekh",
		"Likesh", "Lochan", "Lohit","Lohitashwa", "Lohith", "Lokbhushan", "Lokesh", "Loknath", "Lokpradeep", "Lokprakash", "Lord Shiva", "Lukesh", "Luvya","Maan",
		"Madan","Madangopal", "Madanmohan", "Madesh", "Madhav", "Madhav Krishna", "Madhu",
		"Madhughosh", "Madhujit", "Madhukant", "Madhukar", "Madhup", "Madhusudhana", "Madhusudhana Krishna", "Mahabahu", "Mahabala", "Mahabali", "Mahaddev", "Mahadev", "Mahamati", "Maharanth", "Mahasvin", "Mahavir",
		 "Maheepati", "Mahendra", "Mahesh", "Maheshkumar", "Maheshwar", "Mahipal", "Mahipati", "Mahir",
		"Mahish", "Maitreya", "Maitreya", "Makarand", "Malhar", "Manavendra", "Mandeep", "Mandhatri", "Maneet", "Manendra", "Mangal", "Mangesh", "Mani",
		"Manibhushan", "Manidhar", "Manik","Manikandan", "Manindra", "Maniram", "Manish", "Manishankar", "Manivannan", "Manjeet", "Manjughosh", "Manjunath", "Manmohan", "Manohar", "Manoj",
		"Manoop", "Manoranjan", "Manpal", "Manprasad",
	
}



indianFemaleFirstNames = []string{
	"Aadhira", "Aahna", "Aakaanksha", "Aaloka", "Aanandhi", "Aaral", "Aaruthira", "Aaryahi", "Aashi", "Aashirya", "Aashna", "Aasmi", "Aavani", "Abha",
	"Abhibhushpam", "Abhilasha", "Abhinandana", "Abiradhi", "Achala", "Achiraprabha", "Adah",
	"Adarsha", "Adharshana", "Adhika", "Adhira", "Adhita", "Adhithyaprabha", "Adisakthi", "Aditi", "Adrija", "Adrika", "Advika", "Agalya", "Agasti", "Agnaeyie", "Agnikaa", "Agnimithra", "Agnimitra", "Agniprava", "Agrani", "Ahalya", "Ahana", "Aisha", "Aishani", "Aishwarya", "Ajalaa", "Ajanta", "Akalka", "Akhila", "Akriti", "Akshita", "Akta",
	"Akuti", "Alisha", "Alka",
	"Alpana", "Amaravathy", "Amari", "Amaya", "Amba",
	"Amber", "Ambika", "Ambuja", "Amita", "Amla",
	"Amodini", "Amoli", "Amolika", "Amrapali", "Amrita", "Amrusha", "Amulya", "Anahita", "Anamika", "Anandi", "Anandini", "Ananya", "Anasuya, Anasooya", "Anaya", "Anchal", "Angana", "Anganaa", "Angarika", "Angee", "Anika", "Anindita", "Anisha", "Anishaa", "Anishka", "Anita", "Anjali", "Anjana", "Anju",
	"Anjushree", "Ankita", "Anmol", "Annapurna", "Anshika", "Anshu", "Anshula", "Antara", "Anu",
	"Anubha", "Anuhya", "Anuja", "Anuka", "Anukrithi", "Anumati", "Anumega", "Anumita", "Anupa", "Anupama", "Anupriya", "Anura", "Anuradha", "Anurupika", "Anusha", "Anushka", "Anushri", "Anusuya", "Anuttara", "Anvi",
	"Anvita", "Aparijita", "Aparna", "Apeksha", "Apinaya", "Apsara", "Apurva, Apoorva", "Aradhana", "Arati, Aaarti, Arti", "Archa", "Archan", "Archana", "Ardra", "Arpana", "Arpita", "Arshia", "Aru",
	"Aruna", "Arundhati", "Arunima", "Arushi", "Arya",
	"Aryahi", "Aseema, Ashima", "Asha",
	"Ashakiran", "Ashira", "Ashita", "Ashmita", "Ashna", "Ashnaa", "Ashni", "Ashwini", "Asita", "Asmee", "Asmita", "Astha", "Athalia", "Atreyi", "Aushi", "Avani", "Avanika", "Avanti", "Avantika", "Avika", "Avishi", "Avnita", "Avnitha", "Ayesha", "Ayukta", "Azmin", "Baani", "Badhra", "Badriya", "Bagyawathi", "Bahula", "Bairavi", "Bakula", "Bala",
	"Bandana", "Bandhura", "Banhi", "Banita", "Bansuri", "Barkha", "Baruna", "Baruni", "Basanti", "Bela",
	"Bhaarati", "Bhadirathi", "Bhadra", "Bhagavathi", "Bhagirathi", "Bhagwanti", "Bhagyalakshmi", "Bhagyashree", "Bhagyashri", "Bhairavi", "Bhakti", "Bhamini", "Bhandhavi Bhanu", "Bhandhavya", "Bhanumathi", "Bhanupriya", "Bharati", "Bhargavi", "Bhavana , Bhavna", "Bhavani", "Bhavika", "Bhavini", "Bhavya", "Bhilangana", "Bhooma", "Bhoomika", "Bhuvana", "Bijal", "Bijli", "Bindiya, Bindu", "Bindu", "Bipasha", "Bishakha", "Brinda", "Bulbul", "Cauvery, Cavery", "Chaaya", "Chaitali", "Chaitali, Chaitalee", "Chaitra", "Chameli", "Champa", "Chanchal", "Chanda", "Chandana", "Chandani, Chandini", "Chandanika", "Chandhini", "Chandika", "Chandini", "Chandra", "Chandrakala", "Chandrakanta", "Chandrani", "Chandraprabha", "Chandraswaroopa", "Chandrima", "Charita", "Charitrya", "Charu", "Charulata", "Charulekha", "Charunetra", "Charusmita", "Charvi", "Cheshta", "Chetana", "Chetna", "Chhavi", "Chhaya", "Chintana", "Chintanika", "Chitra", "Chitradevi", "Chitralekha", "Daksha", "Dakshata", "Dakshi", "Damayanti", "Damini", "Darpan", "Darpana", "Darpita", "Darsha", "Darshini", "Daya",
	"Dayamayee", "Dayanita", "Deekshaa", "Deepa", "Deepali", "Deepashikha", "Deepika", "Deepti", "Devak", "Devaki", "Devanee", "Devangana", "Devangi", "Devashree", "Devayani", "Deveshi", "Devi",
	"Devika", "Devmani", "Devyani", "Dhanashri", "Dhanishta", "Dhanista", "Dhanya", "Dharini", "Dharmi", "Dhiraja", "Dhriti", "Dhwani", "Diksha", "Dipali", "Dipashri", "Dipti", "Dipu",
	"Disha", "Dishita", "Diva",
	"Divya", "Drishti", "Drishya", "Dristi", "Dulari", "Durga", "Dwaitha", "Eesha", "Eeshika", "Eiravati", "Ekanthika", "Ektaa, Ekta", "Ela",
	"Elakshi", "Elamathi", "Elina", "Eravati", "Erisha", "Esha",
	"Esha",
	"Eshana", "Eshani", "Eshita", "Eva",
	"Ezhil", "Falak", "Falguni", "Falguni, Phalguni", "Farha", "Faria", "Firaki", "Gaandhaari", "Gaargee", "Gaathaa", "Gaayatree", "Gajagamini", "Gajalakshmi", "Gajlakshmee", "Gajra", "Gandha", "Gandhaa", "Gandhali", "Gandhari", "Ganga", "Gangi", "Gangika", "Gangotri", "Gargi", "Garima", "Gathika", "Gaura", "Gauri", "Gaurika", "Gautami", "Gayatri", "Geena", "Geeti", "Gehna", "Giribala", "Girija", "Gita, Geeta", "Gitanjali", "Gitika", "Godavari", "Gomati", "Gopika", "Gourangi", "Govindini", "Gunasundari", "Gunavati", "Gunjan", "Gunjita", "Gyanda", "Hanita", "Hansa", "Hansika", "Hansini", "Harini", "Haripriya", "Harishita", "Harita", "Harsh", "Harsha", "Harshika", "Harshini", "Harshita", "Harshita", "Hasanthi", "Heena", "Heer",
	"Heera", "Hema",
	"Hemali", "Hemanga", "Hemangi", "Hemkanta", "Hena",
	"Henna", "Himani", "Hiral", "Hiranya", "Hiya",
	"Idhika", "Iha",
	"Ila",
	"Ina",
	"Inaaya", "Indira", "Indrakanta", "Indrakshi", "Indrani", "Indrayani", "Indrina", "Indu",
	"Indu",
	"Induja", "Induja", "Indukala", "Indukanta", "Indumati", "Indumukhi", "Induprabha", "Inrani", "Ipsa",
	"Ipsita", "Ira",
	"Iram",
	"Iravati", "Irit",
	"Isha",
	"Ishana", "Ishani", "Ishanvi", "Ishita", "Ishta", "Ishwari", "Ishwari", "Ishwarya", "Iti",
	"Iyla",
	"Jagriti", "Jahnavi", "Jalaja", "Jamini", "Jamuna", "Janaki", "Janani", "Janhavi", "Jasmin", "Jaya",
	"Jayani", "Jayanti", "Jayaprabha", "Jayaprada", "Jayashri", "Jayasudha", "Jayita", "Jeevana", "Jeevika", "Jhalak", "Jhanvi", "Jharna", "Jheel", "Jhilmil", "Jigyasa", "Jivanta", "Jivika", "Jiya",
	"Joshika", "Joshita", "Juhi",
	"Juily", "Jyoti", "Jyotika", "Jyotsna", "Jyotsnapriya", "Kaanchana", "Kaashi", "Kaavya", "Kadambari", "Kadambini", "Kaishori", "Kaivalya", "Kajal, Kajol", "Kajjali", "Kajri", "Kala",
	"Kalavati", "Kali",
	"Kalika", "Kalindi", "Kalka", "Kallola", "Kallolee", "Kalpana", "Kalpini", "Kalpita", "kalpitha", "Kalya", "Kalyani", "Kamakshi", "Kamakya", "Kamala", "Kamalakshi", "Kamalika", "Kamalini", "Kamana", "Kamika", "Kamini", "Kamna", "Kamnika", "Kamya", "Kanak, Kanaka", "Kanakapriya", "Kanakpriya", "Kanan", "Kananbala", "Kanchan", "Kanchi", "Kaneera", "Kanika", "Kanishka", "Kanta", "kanthi", "Kanti", "Kanupriya", "Kapila", "Kareena", "Karishma", "Karthika", "Kartyayani", "Karuna", "Kashi", "Kashika", "Kashish", "Kashmira", "Kashvi", "Kashyapi", "Kasturi", "Kataksha", "Katyayani", "Kausalya", "Kaushalya", "Kausudhi", "Kaveri", "Kavika", "Kavita", "Keertana", "keerthi", "Kekala", "Kelsy", "Kerani", "Kesari", "Keshi", "Keshika", "Keshini", "Ketaki", "Ketana", "Keya",
	"Khushboo", "Khushi", "Khyati", "Kiara", "Kimaayra", "Kimaya", "kinkini", "Kiran", "Kirti", "Kishori", "Kiya",
	"Kokila", "Komal", "Kopal", "Kopal", "Koyal", "Kranti", "Kripa", "Krishnaveni", "KrishnaVrinda", "Krithi", "kriti", "Kritika", "Krittika", "Kruthika", "Kshanika", "Kshema", "Kshithi", "Kshitija", "Kulina", "Kumari", "Kumkum", "Kumud", "Kumudini", "Kuntal", "Kunti", "Kushali", "Kushboo", "Kusum", "Kusumita", "Kusumitha", "Lajni", "Laksha", "Lakshaki", "Lakshmi, Laxmi", "Lakshmishree", "Lalita, Lalitha", "Lana",
	"Lata",
	"Lavani", "Lavanya", "Leela", "Leelawati", "Leena", "Lekha", "Likhitha", "Lily",
	"Lipi",
	"Lipika", "Loganayaki", "Lola",
	"Lopa",
	"Lopamudra", "Lorena", "Loukya", "Maanika", "Madhavi", "Madhu", "Madhubala", "Madhuksara", "Madhulata", "Madhulika", "Madhumalati", "Madhumati", "Madhumita", "Madhuparna", "Madhur", "Madhuri", "Madhurima", "Madhushri", "Magana", "Mahadevi", "Mahak", "Mahalakshmi", "Mahasweta", "Maheshi", "Mahi",
	"Mahika", "Mahima", "Mahitha", "Maina", "Maithili, Mythili", "Maitri", "Makshi", "Mala",
	"Malati", "Malavika", "Malini", "Mallika", "Manali", "Manana", "Manasa, Maanasa", "Manasi", "Manasi", "Mandakini", "Mandira", "Mandodari", "Mangala", "Manideepa", "Manimala", "Manisha", "Manitha", "Manjari", "Manjika", "Manjira", "Manju", "Manjula", "Manjusha", "Manjusri", "Manmayi", "Manorama", "Manushri", "Manvi", "Matangi", "Maushmi", "Mausumi, Maushmi", "Maya",
	"Mayuri", "Medha", "Meena", "Meenakshi", "Meera", "Megha", "Meghna", "Menka", "Mili",
	"Minal, Meenal", "Minati", "Minaxi", "Mini",
	"Mira",
	"Misha,Meesha", "Mishti", "Mita, Meeta", "Mithilesh", "Mitul", "Mohini", "Mohisha", "Mohita", "Mona",
	"Monica", "Mridula", "Mriganayani", "Mriganka", "Mrinalini", "Mrityunjay", "Mudita", "Mudra", "Mudrika", "Mugdha", "Mukta", "Mukti", "Muskaan", "Mythili", "Naayantara", "Nabhi", "Nagi",
	"Naina", "Naini", "Nainika", "Naishadha", "Nakshatra", "Nalika", "Nalini", "Namita", "Namrata", "Namya", "Nanda", "Nandini", "Nandita", "Naomi", "Narayani", "Narmada", "Natasha", "Navaneeta", "Naveena", "Nayana", "Neela", "Neelabja", "Neelakshi", "Neelam", "Neelambari", "Neelanjana", "Neera", "Neeraja", "Neeru", "Neeta", "Neeti", "Neha",
	"Nehal", "Netra", "Neysa", "Nidhi", "Niharika", "Nihira", "Nikhita", "Nikki", "Nila",
	"Nima",
	"Nimeesha", "Nina, Neena", "Niranjana", "Nirmala", "Nirmayi", "Nirupama", "Nisha", "Nishi", "Nishikanta", "Nishita", "Nishithini", "Nishtha", "Niti",
	"Nitya", "Nityapriya", "Nivedita", "Niyati", "Noopur", "Noori", "Nutan", "Nyna",
	"Ojal",
	"Ojasvi", "Ojaswini", "Olena", "Omisha", "Omisha", "Omkareshwari", "Omkareshwari", "Onella", "Oni",
	"Oorja", "Orpita", "Padimni", "Padma", "Padmaja", "Padmakalyani", "Padmakshi", "Padmal", "Padmini", "Pakhi", "Palak", "Pallavi", "Pallavini", "Panchali", "Pankaja", "Pankhuri", "Panna", "Paramita", "Paramjyothi", "Parasmani", "Paravi", "Pari",
	"Paridhi", "Parijat", "Parina", "Parinaaz", "Parineetha", "Parni", "Parnika", "Paru",
	"Parul", "Parvani", "Parvati", "Parveen", "Pauravi", "Pavana", "Pavani", "Pavi",
	"Pavithritha", "Pavitra", "Payal", "Phalak", "Phalguni", "Pia",
	"Piyali", "Pooja, Puja", "Poojitha", "Poonam", "Poorbi", "Poorineeta", "Poorna", "Poorva", "Poorvaganga", "Prabha", "Prabharoopa", "Prabhati", "Prabhrithi", "Prachi", "Pradeepa", "Pradeepta", "Pradhi", "Pradhyumna", "Pragalbha", "Pragati", "Pragya", "Prajna", "Prakriti", "Prakruti", "Pramada", "Pramila", "Pramiti", "Pramuditha", "Pranaali", "Pranati", "Prapthi", "Prapti", "Prarthana", "Praseeda", "Prashansa", "Prashanti", "Prasutha", "Prateeksha", "Prathitha", "Prathysha", "Pratibha", "Pratima", "Pratishtha", "Pratyaya", "Pratyusha", "Prayaathi", "Preetha", "Preeti", "Prem",
	"Prema", "Prerana", "Prerita", "Preyasi", "Prisha", "Prita", "Priti", "Priya", "Priyadarshini", "Priyal", "Priyam", "Priyamvada", "Priyanka", "Priyatha", "Pujita", "Pujya", "Pulakitha", "Punita", "Punya", "Purnima, Poornima", "Purva", "Purvaja", "Pushpa", "Pushpanjali", "Pushpita", "Pushpitha", "Pushthi", "Raakhi, Rakhi", "Rachana", "Radha", "Radhika", "Ragini", "Rajalakshmi", "Rajani, Rajini", "Rajanigandha", "Rajasi", "Rajeshvari", "Raji",
	"Rajkumari", "Rajshri", "Rajvi", "Raka",
	"Rakhi", "Raksha", "Raktima", "Ramani", "Ramanika", "Rambha", "Ramita", "Ramona", "Ramya", "Rani",
	"Ranita", "Ranjana", "Ranjini", "Ranjita", "Rashi, Raashi", "Rashmi", "Rashmika", "Rasika", "Rasna", "Rati",
	"Ratna", "Ratnali", "Rea, Ria", "Reetul", "Rekha", "Rena",
	"Renu",
	"Renuka", "Resham", "Reshma", "Reshmi", "Resna", "Reva",
	"Revati", "Rewa",
	"Rhea",
	"Richa", "Riddhi", "Rimjhim", "Rishika", "Rishima", "Rishita", "Ritika", "Ritu",
	"Riya",
	"Rohini", "Rohita", "Roma",
	"Romila", "Roopam", "Roshana", "Roshni",
}

indianLastNames = []string{
	"abbi", "abhyankar", "abraham", "acharya", "achrekar", "adani", "adhikane", "adhikari", "adiga", "advani", "agarkar", "agarwal", "agate", "agnihotri", "agvan", "ahlawat", "ahluwalia", "ahuja", "aiyappa", "ajgaonkar", "akash", "akhtar", "akkineni", "ali", "amarnath", "ambani", "ambedkar", "ambekar", "amble", "ambuja", "ambulkar", "ameer", "amin", "amladi", "amravatikar", "amrutham", "anand", "anandi", "anchan", "anchan", "aneja", "annaladasula", "antony", "anumula", "apdev", "apte", "apture", "argade", "arjun", "arora", "arya", "ashar", "asrani", "asthana", "ashtikar", "atre", "atrey", "atri", "attal", "atwal", "aulakh", "avari", "awasthi", "babariya", "babbar", "babu", "bachchan", "badgujar", "badkas", "badhe", "bahl", "bhadsavle", "bhaduri", "bhasi", "bafna", "baghel, bagul", "baid", "bajaj", "bajpai", "bajwa", "bakshi", "balaji", "balakrishnan", "balasubramaniam", "balsara", "bamnote", "bandodkar", "bandyopadhyay", "banerjee", "bangera", "bangre", "banjan", "bansal", "bansod", "banthia", "banthiya", "bantwal", "bapat", "barde", "bardia", "barman", "barot", "barua", "basra", "basu", "batliwalla", "batra", "basha", "bawaskar", 
	"baweja", "bawre", "bedi", "bendre", "benegal", "bengre", "berde", "bhaiya", "bhagat", "bhagavathar", "bhagwat", "bhakri", "bhalla", "bhalodiya", "bhan", "bhandari", "bhanot", "bhanushali", "bharadwaj", "bhardwaj", "bhargav", "bharjatya", "bharti", "bhat", "bhathheja", "bhatia", "bhattacharya", "bhattad", "bhatti", "bhave", "bhavsar", "bhende", "bhide", "bhimani", "bhite", "bhoir", "bhoite", "bhojwani", "bhosale", "bhowmick", "bhupathi", "bhushan", "bhutada", "bijlani", "bika", "binny", "birla", "bisht", "biswas", "biyani", "bobde", "bodhankar", "bokare", "bora", "borgaonkar", "borkar", "bose", "bradoo", "buddharaju", "budhisagar", "bundela", "burman", "busanagari", "buty", "cethirakath", "chabbria", "chabukswar", "chacko", "chahal", "chakole", "chakraborty", "chakravarthy", "chanana", "chanchad", "chandani", "chandavarkar", "chandel", "chander",
	 "chandak", "chandok", "chandorkar", "chandrababu", "channe", "chafekar", "chaphekar", "chary", "chattaraj", "chatterjee", "chattopadhyay", "chaturvedi", "chandra", "chaubey", "chauhan", "chavan", "chaurasiya", "chaurasia", "chawla", "chedge", "cheema", "chettipally", "chimalwar", "chimote", "chimurkar", "chintawar", "chitale", "chitalia", "chitnis", "choksi", "chopra", "chorghade", "chotai", "choudhary", "chaudhary", "choudhury", "chowdhury", "chowdhary", "chaudhari", "chugh", "dabholkar", "dabral", "dadlani", "daftari", "daga", "dahake", "dalvi", "damble", "damle", "dandale", "dandekar", "dang", "dani", "darbari", "darji", "darling", "das", "dasgupta", "dasgupta", "dashmunshi", "daswani", "datey", "dave", "dawande", "dayal", "de", "dedhe", "dedhia", "dehadrai", "denzongpa", "deodhar", "deol", "deora", "desai", "deshmukh", "deshpande", "dev", "devadiga", "devadikar", "devaiah", "devgade", "devgan", "dewaikar", "dey", "dhake",
	  "dhaliwal", "dhameja", "dhanoa", "dhar", "dharashiokar", "dharmadhikari", "dhawan", "dhingra", "dholakia", "dhone", "dhoni", "dhoot", "dhote", "dhumal", "dhumale", "dikshit", "divekar", "diwadkar", "diwe", "dixit", "dobriyal", "doifode, doiphode", "dongre", "dosanjh", "doshi", "draboo", "dua", "dubey", "duddala", "dugar", "duggal", "dusanj", "dutt", "dutta", "dwivedi", "eknath", "engle, ingle", "faasil", "fadikar", "fadnavis", "fansekar", "faujdar", "faye", "fernandez", "firodia", "fotedar", "gadde", "gadekar", "gadgil", "ghadse", "gadikar", "gadkari", "gadre", "gaikwad", "gaiki", "gaitonde", "gajjar", "gambhir", "ganatra", "gandhi", "ganesan", "gangopadhyay", "gangotri", "ganguly", "garach", "garapati", "garcia", "garg", "garware", "gavaskar", "gaur", "gautam", "gavit", "gawande", "gayatri", "gehlot", "gera", "ghai", "ghaisas", "ghanekar", "ghatate", "ghate", "ghosh", "ghoshal", "ghui", "gilani", "gill", "girdhar", "girotra", "godbole", "godse", "goel", "goenka", "gohad", "gohil", "gokarn", "gokhale", "golani", "gole", "gopal", "gopalan", "gopi",
	 "gore", "gosai", "gosain", "goswami", "gounder", "govind", "govitrikar", "gowarikar", "gowda", "goyal", "grewal", "grover", "guha", "gujar", "gujral", "gulgule", "gulwadi", "gunnam", "gupta", "gupte", "gurnani", "gursahani", "gurwara", "haasan", "hadapsar", "hadas", "hagavane", "haksar", "handoo", "hangal", "hansraj", "hardas", "haridas", "harode", "hashmi", "hasnee", "hassan", "hattangadi", "hazra", "hazare", "hazari", "hazarika", "hebbar", "hegde", "hinduja", "hingorani", "hiranandani", "hirani", "hirlekar", "hirwani", "hiwre", "hole", "holkar", "hooda", "hoon", "hoskote", "huilgol", "hundal", "hussainl", "inamdar", "indoria", "indulkar", "ingle", "ingole", "iragavarapu", "irani", "ivaturi", "iyengar", "iyer", "jadeja", "jadhav", "jadi", "jadon", "jagtap", "jain", "jaiswal", "jajoo", "jalan", "janardhan", "janefalkar", "janmeja", "janolkar", "janumala", "jariwala", "jatasra", "jawalkar", "jena", "jha", "jhathar", "jhawar", 
	 "jindal", "jinturkar", "jobanputra", "joglekar", "johal", "johar", "joshi", "juneja", "juvekar", "kadakia", "kadam", "kak", "kakde", "kalawar", "kalbhor", "kale", "kalra", "kalwani", "kamat", "kambli", "kamboj", "kanakia", "kanchan", "kane", "kanetkar", "kanhe", "kanitkar", "kankariya", "kannaujiya", "kanojiya", "kanoongo", "kansal", "kapadia", "kapoor", "kapre", "kapse", "kapur", "kar", "karanth", "karia", "karkare", "karkera", "karlekar", "karnad", "karnatakkachu", "karnawat", "karnik", "karthik", "karwande", "kasat", "kasundra", "katoch", "kaul", "kazi", "kazmi", "kelkar", "keshari", "kewalramani", "khadse", "khaitan", "khan", "khandar", "khandelwal", "khanduja", "khanna", "khanolkar", "kharbanda", "khatri", "khedekar", "khemu", "kher", "khetarpal", "khobragade", "kholkute", "khopkar", "khosla", "khurana", "kilachand", "kinariwala", "kinariwalla", "kirloskar", "kochar", "koganti", "kohale", "kondke", "kothale", "kothari",
	  "kotian", "kottarakkara", "kottary", "kriplani", "krishna", "krishnamurthy", "kuchibhotla", "kukde", "kukkar", "kukreja", 
	 "kukyan", "kulkarni", "kulshreshtha", "kumari", "kumbhare", "kumble", "kumpawat", "kumta", "kundar", "kundra, kunder", "kurien", "kurup", "kathiriya", "koshley", "lad", "laddha", "lagadapati", "laghate", "lagoo", "lahoti", "lakhani", "lakhotia", "lal", "lalbhai", "lalchandani", "lallad", "lalwani", "lamba", "langote", "langroo", "lapalikar", "latkar", "lele", "ligam", "limaye", "lobo", "lodha", "lohia", "lokhande", "lopes", "lote", "ludhani", "lunkad", "luthra", "maan", "mabiyan", "madasani", "madiwale", "mafatlal", "mahajan", "mahakale", "mahalingam", "mahashabde", "maheshwari", "mahindra", "mahindrakar", "maitra", "majumdar", "makhija", "makwana", "malhan", "malhotra", "malhotra", "mali", "malik", "malini", "mallapur", "mallya", "malpe", "malusare", "malviya", "mamdani", "manak", "manchandha", "mandal", "mandalik", "mande", "mandhane", "mandlekar", "manjrekar", "marathe", "marwah", "marwaha", "mashelkar", "mathur",
	  "mathurkar", "mattu", "mayar", "meghe", "mehra", "mehrotra", "mehrotra", "mehta", "mendhi", "mendon", "menon", "merchant", "mhatre", "minhas", "mirchandani", "mishra", "mistry", "mittal", "mittal", "modak", "modi", "moghe", "mohan", "mohanty", "mohit", "mohril", "moitra", "mojala", "molkar", "mongia", "more", "morea", "motwani", "motwani", "macwan", "mudaliyar", "mujumdar", "mukherjee", "mulchandani", "mule", "muley", "mulki", "mundan", "mundhe", "mundi", "munjal", "munje", "munshi", "muralidharan", "murthy", "mushrif", "muttemwar", "muzumdar", "noora", "nabar", "nadar", "nadella", "nadikatla", "nadkarni", "nag", "nagaiah", "nagar", "nagpal", "naik", "nair", "nakra", "nalamolu", "nambian", "nambiar", "nambiyar", "namboodiri", "nakhate", "nanavati", "nanda", "nandamuri", "nandanwar", "nandigam", "narang", "narayan", "narayanan", "naresh", "narlikar", "narula", "nath", "nathawat", "naudiyal", "nawre", "nayak", "negi", 
	"nehra", "nehru", "nelluri", "nene", "nerurkar", "nigam", "nihalani", "nikam", "nike", "nikhanj", "nimbalkar", "nimhan", "nischol", "nitharwal", 
	 "oak", "oberoi", "ogle, ogale", "ogra", "ohri", "omble", "omkareshwar", "omkarnath", "ozha", "padukone", "padwad", "pahwa", "pai", "pagi", "paintal", "pal", "palan", "palande", "palandurkar", "palav", "palekar", "paliwal", "palod", "pancholi", "panda", "pande", "pandey", "pandher", "pandhripande", "pandian", "pandit", "pandya", "pangarkar", "page", "pagey", "panicker", "panigrahi", "panjwani", "panse", "pant", "panthulu", "panwar", "pappu", "papule", "parab", "paranjape", "parasher", "parate", "parchure", "pardesi", "pardeshi", "parekh", "parihar", "parikh", "parkhi", "parmar", "parsodkar", "paruchuri", "pasricha", "patankar", "patekar", "patel", "pathak", "patil", "patki", "patnaik", "patne", "patni", "patra", "patwardhan", "paud", "paul", "pavri, pawri", "pawar", "pednekar", "pendharkar", "pendse", "phanse", "phatak", "phogat", "pillai", "pimparkar", "pimplapure", "pingale", "pinjarkar", "piramal", "piramal", "pitale",
	  "pohankar", "pohnekar", "polekar", "porwal", "potdar", "potdukhe", "potluri", "prabhavalkar", "prabhu", "pradhan", "prajapati", "prasad", "prasade", "prasadi", "pratham", "premji", "punia", "puniani", "puniyani", "punj", "punwani", "puranik", "puri", "purohit", "puthran", "prakash", "raaju", "raaz", "rahane", "raheja", "rahman", "rai", "raikantiwar", "raina", "raizada", "raj", "rajan", "rajawat", "rajpurohit", "raju", "rakshit", "ramachandran", "ramelwar", "ramesh", "ramisetty", "ramnani", "rampal", "rana", "ranade", "ranawat", "rane", "rao", "rastogi", "rathi", "rathod", "rathore", "ravinuthala", "rawal", "rawat", "rawat", "raxit", "ray", "raykantiwar", "reddy", "rege", "rehman", "reshammiya", "rode", "rokade", "roshan", "roy", "roychoudhury", "ruia", "rungta", "ruparel", "sabnis", "sachan", "sachdev", "sadarangani", "sadhu", "safary", "saha", "sahai", "sahastrabuddhe", "sahedev", "sahni", "sahu", "sahukar", "saigal", "saini", "saklani", "sakpal", "salian",
	 "salmaan", "saluja", "salunkhe", "samarth", "samdurkar", "samudre", "sandhu", "sanghavi", "sanghvi", "sanil", "sanir", "sanon", "saoji", "sapra", "saraf", "sarda", "sarja", "sarnaik", "sarode", "sarve", "sarwe", "satam", "satija", "savarkar","sawant","sawarkar", "sawhney", "sawji", "saxena", "sehgal", "sehwag", "sekhri", "selvan", "sen", "sengupta", "sethi", "setna", "shah", "shahane", "shareef", "sharma", "shastri", "shekhawat", "shelke", "shenoy", "sheikh", "sheirgill", "shekhar", "sheshadri", "shete", "sheth", "shetty", "shikhavat", "shikhawat", "shinde", "shirke", "shishodia", "shivalkar", "shivdasani", "shourey", "shrikhande", "shriyan", "shroff", "shukla", "siddiqui", "sidhu", "sikarwar", "[simha]]", "singh", "singhal", "singham", "singhania", "sinha", "siripurapu", "sirish", "sobti", "solanki", "somaiya", "somalwar", "soman", "somani", "somayajulu", "sonawane", "soni", "sonolikar", "sonowal", "sonpatki", "sood", 
	 "srinivasan", "srivastava", "subramaniam", "subramanium", "subramanyam", "suchak", "suman", "surana", "suri", "surve", "suvarna", "swamy", "swarup", "syed", "tadaskar", "tagore", "tahil", "tahiliani", "tahilramani", "tak", "talari", "talgeri", "talpade", "talsania", "talwar", "tambe", "tambke", "tamang", "tamhane", "tamhankar", "tandel", "tandon", "tanna", "tanti", "tanwar", "taparia", "tata", "taurani", "tawde", "tayde", "tejwani", "tekchandani", "tempalli", "tendulkar", "thackeray", "thakersey", "thakkar", "thakore", "thakral", "thakran", "thakre", "thakrey", "thakur", "thakurtha", "thapa", "thapar", "thombre", "thorvi", "thosar", "thota", "thuse", "tijori", "tikoo", "tilak", "tirodkar", "tirpude", "tiwari", "tiwaskar", "tomar", "toor", "tope", "toshniwal", "tripathi", "trivedi", "tufchi", "tyagi", "tapaniya", "udayaraju", "udyavar", "ullal", "umbarkar", "unhale", "uniyal", "unni", "unnikrishnan", "unnithan", "upadhyay", 
	 "upadhye", "upasni", "uplenchwar", "uppal", "uppalapati", "upponi", "velugubanti", "vaghela", "vaida", "vaidya", "vaish", "vajpayee",
	 "valecha", "varghese", "varma", "vartak", "vashisth", "vasudevan", "vaswani", "vaze", "veeranna", "velankar", "vengsarkar", "venkat", "venkateswaran", "venkatraman", "verghese", "verma", "vernekar", "vichare", "vidyarthi", "vij", "vijay", "vijayrania", "vincent", "virani", "virk", "virmani", "visariya", "vora", "vyas", "wadekar", "wadhavkar", "wadhva", "wadhawa", "wadhawan", "wadhwani", "wadia", "wagh", "waghe", "waghmare", "waghray", "wagle", "waingankar", "wajantri", "waknis", "walale", "walawalakar", "walia", "walker", "wangdu", "wangnoo", "wani", "wanjare", "wanjari", "wankar", "wankhede", "waradpande", "wargantiwar", "warhadpande", "warraich", "warsi", "watane", "watharkar", "wattal", "watwe", "wazalwar", "wazir", "yadav", "yagnik", "yarlagadda", "yedekar", "yederi", "yelimeli", "yelkar", "yelpude", "yenkie", "yennemadi", "yenugula", "yeolekar", "yerukola", "yesuraju", "zaantye", "zade", "zahaldar", "zakaria", "zalpuri", 
	 "zarapkar", "zariwala", "zate", "zaveri", "zende", "zite", "zutshi", "tadaskar", "tagore", "tahil", "tahiliani", "tahilramani", "tak", "talari", "talgeri", "talpade", "talsania", "talwar", "tambe", "tambke", "tamang", "tamhane", "tamhankar", "tandel", "tandon", "tanna", "tanti", "tanwar", "taparia", "tata", "taurani", "tawde", "tayde", "tejwani", "tekchandani", "tempalli", "tendulkar", "thackeray", "thakersey", "thakkar", "thakore", "thakral", "thakran", "thakre", "thakrey", "thakur", "thakurtha", "thapa", "thapar", "thombre", "thorvi", "thosar", "thota", "thuse", "tijori", "tikoo", "tilak", "tirodkar", "tirpude", "tiwari", "tiwaskar", "tomar", "toor", "tope", "toshniwal", "tripathi", "trivedi", "tufchi", "tyagi", "tapaniya", "udayaraju", "udyavar", "ullal", "umbarkar", "unhale", "uniyal", "unni", "unnikrishnan", "unnithan", "upadhyay", "upadhye", "upasni", "uplenchwar", "uppal", "uppalapati", "upponi", "velugubanti",
	  "vaghela", "vaida", "vaidya", "vaish", "vajpayee", "valecha", "varghese", "varma", "vartak", "vashisth", "vasudevan", "vaswani",
	 "vaze", "veeranna", "velankar", "vengsarkar", "venkat", "venkateswaran", "venkatraman", "verghese", "verma", "vernekar", "vichare", "vidyarthi", "vij", "vijay", "vijayrania", "vincent", "virani", "virk", "virmani", "visariya", "vora", "vyas", "wadekar", "wadhavkar", "wadhva", "wadhawa", "wadhawan", "wadhwani", "wadia", "wagh", "waghe", "waghmare", "waghray", "wagle", "waingankar", "wajantri", "waknis", "walale", "walawalakar", "walia", "walker", "wangdu", "wangnoo", "wani", "wanjare", "wanjari", "wankar", "wankhede", "waradpande", "wargantiwar", "warhadpande", "warraich", "warsi", "watane", "watharkar", "wattal", "watwe", "wazalwar", "wazir", "yadav", "yagnik", "yarlagadda", "yedekar", "yederi", "yelimeli", "yelkar", "yelpude", "yenkie", "yennemadi", "yenugula", "yeolekar", "yerukola", "yesuraju", "zaantye", "zade", "zahaldar", "zakaria", "zalpuri", "zarapkar", "zariwala", "zate", "zaveri", "zende", "zite", "zutshi",
}
)
