package fakeit

// All word lists and lookup tables used by Generator methods.

var firstNames = []string{
	"Alice", "Bob", "Carol", "David", "Eve", "Frank", "Grace", "Hank",
	"Iris", "Jack", "Karen", "Liam", "Mia", "Noah", "Olivia", "Peter",
	"Quinn", "Rachel", "Sam", "Tara", "Uma", "Victor", "Wendy", "Xander",
	"Yara", "Zoe", "Aaron", "Bella", "Carlos", "Diana", "Ethan", "Fiona",
	"George", "Hannah", "Ivan", "Julia", "Kevin", "Laura", "Marcus", "Nina",
	"Oscar", "Penelope", "Quentin", "Rosa", "Stefan", "Theresa", "Ursula",
	"Vincent", "Willa", "Xavier", "Yvonne", "Zachary",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller",
	"Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez",
	"Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
	"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark",
	"Ramirez", "Lewis", "Robinson", "Walker", "Young", "Allen", "King",
	"Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores", "Green",
	"Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell",
	"Carter", "Roberts",
}

var adjectives = []string{
	"quick", "lazy", "bright", "dark", "silent", "loud", "cold", "warm",
	"ancient", "modern", "hidden", "open", "sharp", "smooth", "rough",
	"simple", "complex", "bold", "gentle", "wild", "tame", "tiny", "giant",
	"brave", "clever", "fierce", "calm", "dusty", "golden", "silver",
	"hollow", "mystic", "frozen", "blazing", "drifting", "rustic", "vibrant",
	"noble", "humble", "proud", "swift", "steady", "wandering", "lost",
	"found", "restless", "quiet", "curious", "fearless",
}

var nouns = []string{
	"river", "mountain", "forest", "sky", "ocean", "desert", "city",
	"village", "tower", "bridge", "stone", "flame", "wind", "shadow",
	"light", "cloud", "storm", "wave", "valley", "peak", "path", "gate",
	"field", "harbor", "island", "lake", "shore", "cave", "cliff", "mesa",
	"plain", "grove", "hollow", "crest", "ridge", "basin", "delta", "reef",
	"inlet", "dune", "moor", "fen", "glade", "dell", "knoll", "butte",
	"fjord", "steppe", "tundra", "savanna",
}

var animals = []string{
	"fox", "wolf", "bear", "eagle", "hawk", "owl", "tiger", "lion",
	"deer", "rabbit", "otter", "beaver", "falcon", "raven", "crane",
	"viper", "cobra", "panther", "lynx", "moose", "elk", "bison", "whale",
	"dolphin", "shark", "seal", "walrus", "puffin", "heron", "kite",
}

var emailDomains = []string{
	"gmail.com", "yahoo.com", "outlook.com", "proton.me", "icloud.com",
	"fastmail.com", "hey.com", "tutanota.com", "zoho.com", "pm.me",
}

var domains = []string{
	"example.com", "techcorp.io", "devhub.net", "cloudsync.ai",
	"opendata.org", "saasify.dev", "buildfast.co", "nexusapp.io",
	"gridwork.net", "stacklabs.dev",
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/605.1.15 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Android 14; Mobile; rv:121.0) Gecko/121.0 Firefox/121.0",
	"Go-http-client/2.0",
	"python-requests/2.31.0",
	"curl/8.4.0",
	"PostmanRuntime/7.36.1",
	"axios/1.6.2",
}

var cities = []string{
	"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia",
	"San Antonio", "San Diego", "Dallas", "San Jose", "Austin", "Jacksonville",
	"London", "Paris", "Berlin", "Tokyo", "Seoul", "Mumbai", "São Paulo",
	"Mexico City", "Cairo", "Lagos", "Nairobi", "Sydney", "Melbourne",
	"Toronto", "Vancouver", "Dubai", "Singapore", "Hong Kong", "Shanghai",
	"Beijing", "Bangkok", "Jakarta", "Karachi", "Istanbul", "Moscow",
	"Madrid", "Rome", "Amsterdam", "Brussels", "Vienna", "Warsaw", "Prague",
	"Budapest", "Zurich", "Geneva", "Stockholm", "Oslo", "Helsinki",
}

var countries = []string{
	"United States", "United Kingdom", "Germany", "France", "Japan",
	"Canada", "Australia", "Brazil", "India", "China", "South Korea",
	"Mexico", "Italy", "Spain", "Netherlands", "Sweden", "Norway",
	"Denmark", "Finland", "Switzerland", "Austria", "Poland", "Portugal",
	"Argentina", "Chile", "Colombia", "South Africa", "Nigeria", "Egypt",
	"Turkey", "Saudi Arabia", "UAE", "Israel", "Pakistan", "Bangladesh",
	"Indonesia", "Philippines", "Vietnam", "Thailand", "Malaysia",
	"Singapore", "New Zealand", "Ireland", "Belgium", "Czech Republic",
	"Hungary", "Romania", "Ukraine", "Russia", "Greece",
}

var countryCodes = []string{
	"US", "GB", "DE", "FR", "JP", "CA", "AU", "BR", "IN", "CN",
	"KR", "MX", "IT", "ES", "NL", "SE", "NO", "DK", "FI", "CH",
	"AT", "PL", "PT", "AR", "CL", "CO", "ZA", "NG", "EG", "TR",
	"SA", "AE", "IL", "PK", "BD", "ID", "PH", "VN", "TH", "MY",
	"SG", "NZ", "IE", "BE", "CZ", "HU", "RO", "UA", "RU", "GR",
}

var timeZones = []string{
	"America/New_York", "America/Los_Angeles", "America/Chicago",
	"America/Denver", "America/Toronto", "America/Vancouver",
	"America/Sao_Paulo", "America/Buenos_Aires", "America/Mexico_City",
	"Europe/London", "Europe/Paris", "Europe/Berlin", "Europe/Madrid",
	"Europe/Rome", "Europe/Amsterdam", "Europe/Stockholm", "Europe/Oslo",
	"Europe/Helsinki", "Europe/Warsaw", "Europe/Moscow",
	"Asia/Tokyo", "Asia/Seoul", "Asia/Shanghai", "Asia/Hong_Kong",
	"Asia/Singapore", "Asia/Bangkok", "Asia/Dubai", "Asia/Karachi",
	"Asia/Kolkata", "Asia/Jakarta",
	"Australia/Sydney", "Australia/Melbourne", "Pacific/Auckland",
	"Africa/Cairo", "Africa/Lagos", "Africa/Nairobi",
}

var currencies = []string{
	"USD", "EUR", "GBP", "JPY", "CAD", "AUD", "CHF", "CNY", "SEK",
	"NZD", "MXN", "SGD", "HKD", "NOK", "KRW", "TRY", "RUB", "INR",
	"BRL", "ZAR", "DKK", "PLN", "THB", "IDR", "CZK", "AED", "SAR",
}

var colorNames = []string{
	"red", "green", "blue", "yellow", "orange", "purple", "pink", "brown",
	"black", "white", "gray", "cyan", "magenta", "lime", "maroon", "navy",
	"olive", "teal", "silver", "gold", "coral", "salmon", "indigo", "violet",
	"crimson", "turquoise", "lavender", "beige", "ivory", "charcoal",
}

var companyPrefixes = []string{
	"Apex", "Blue", "Cloud", "Delta", "Edge", "Fast", "Global", "Horizon",
	"Infinite", "Jet", "Key", "Lynx", "Metro", "Nexus", "Open", "Prime",
	"Quantum", "Rapid", "Sky", "Tech", "Ultra", "Vertex", "Wave", "Xcel",
	"Yield", "Zenith", "Alpha", "Beta", "Core", "Deep", "Echo", "Flux",
}

var companySuffixes = []string{
	"Technologies", "Solutions", "Systems", "Group", "Labs", "Works",
	"Industries", "Ventures", "Partners", "Consulting", "Digital",
	"Analytics", "Dynamics", "Networks", "Services", "Innovations",
	"Enterprises", "Capital", "Media", "Data",
}

var jobLevels = []string{
	"Junior", "Senior", "Lead", "Principal", "Staff", "Head of",
	"VP of", "Director of", "Chief", "Associate",
}

var jobRoles = []string{
	"Software Engineer", "Product Manager", "Data Scientist", "Designer",
	"DevOps Engineer", "Security Analyst", "Backend Engineer", "Frontend Engineer",
	"ML Engineer", "Site Reliability Engineer", "Engineering Manager",
	"QA Engineer", "Technical Writer", "Solutions Architect",
	"Platform Engineer", "Data Engineer", "Research Scientist",
	"Business Analyst", "Scrum Master", "Cloud Architect",
}

var industries = []string{
	"Technology", "Finance", "Healthcare", "Education", "Retail",
	"Manufacturing", "Transportation", "Energy", "Real Estate",
	"Media & Entertainment", "Telecommunications", "Agriculture",
	"Aerospace & Defense", "Pharmaceuticals", "Automotive",
	"Food & Beverage", "Construction", "Legal", "Insurance", "Logistics",
}

var fileExtensions = []string{
	".pdf", ".docx", ".xlsx", ".pptx", ".txt", ".csv", ".json", ".xml",
	".yaml", ".toml", ".md", ".png", ".jpg", ".gif", ".svg", ".mp4",
	".mp3", ".zip", ".tar.gz", ".sh", ".py", ".go", ".js", ".ts",
}

var mimeTypes = []string{
	"application/json", "application/xml", "application/pdf",
	"application/zip", "application/octet-stream",
	"text/plain", "text/html", "text/csv", "text/markdown",
	"image/png", "image/jpeg", "image/gif", "image/svg+xml", "image/webp",
	"audio/mpeg", "audio/ogg", "video/mp4", "video/webm",
	"multipart/form-data", "application/x-www-form-urlencoded",
}

var httpMethods = []string{
	"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
}

var httpStatusCodes = []int{
	200, 201, 204, 301, 302, 400, 401, 403, 404, 409, 422, 429, 500, 502, 503,
}

var streetNames = []string{
	"Oak", "Maple", "Cedar", "Pine", "Elm", "Birch", "Walnut", "Chestnut",
	"Main", "Market", "Park", "Lake", "River", "Hill", "Valley",
	"Sunset", "Sunrise", "Spring", "Summer", "Autumn",
	"Lincoln", "Washington", "Jefferson", "Madison", "Monroe",
}

var streetSuffixes = []string{
	"Street", "Avenue", "Boulevard", "Drive", "Lane", "Road", "Way",
	"Court", "Place", "Terrace", "Circle", "Trail", "Path",
}

var loremText = `Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua ut enim ad minim veniam
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur excepteur sint occaecat cupidatat
non proident sunt in culpa qui officia deserunt mollit anim id est laborum`
