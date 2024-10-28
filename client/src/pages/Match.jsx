import ButtonLine from "../components/buttons/ButtonLine"
import ButtonRound from "../components/buttons/ButtonRound"
import CardProfile from "../components/forms/CardProfile"
import HeaderWhite from "../components/headers/HeaderWhite"
import Itinerary from "../components/itineraries/Itinerary"
import Scroll from "../layouts/Scroll"

function Match() {
  const [data, setData] = useState({})
  const [name, setName] = useState()
  const [gender, setGender] = useState()
  const [country, setCountry] = useState()

  const { state } = useLocation();
  const { activity, trip, finderId } = state
  const navigate = useNavigate()
   
  useEffect(() => {
    try {
      const fetchData = async () => {
        try {
          const response = await fetch(
            `http://localhost:8000/api/itinerary/buddy/destination?activity=${activity}&trip=${trip}`
          );
          const data = await response.json();
          setData(data);
          setName(data.user.name)
          setGender(data.user.gender)
          setCountry(data.user.country)
        } catch (error) {
          console.log(error);
        }
      };

      fetchData();
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }, []);
  
  const buttonConfirm = async () => {
    const userId = localStorage.getItem("user")
    let join;
    try {
      join = await fetch("http://localhost:8000/api/itinerary/buddy/join", {method: "POST", body: JSON.stringify({
        itinerary_id: data.itinerary_id,
        itinerary_finder_id: finderId,
        created_by: parseInt(userId),
      }),
      headers: {
        'Content-Type': 'application/json' // Tipe konten JSON
      }});
      join = await join.json()
      console.log("INI", join)
    } catch (error) {
      console.log(error)
    } finally {
      navigate("/home", {state: {joinId: join.ID}});
    }
  };
  
  return (
    <Scroll>
      <div className="flex flex-col items-center w-full h-96 bg-blue">
        <HeaderWhite />

        <CardProfile />

        <div className="relative flex h-fit w-96 flex-col gap-8 rounded-t-3xl bg-white px-6 pt-5 text-sm shadow-[0_-4px_20px_rgba(0,0,0,0.1)]">
          <div className="absolute flex gap-4 -top-6 right-8">
            <ButtonRound />
            
            <ButtonLine />
          </div>

          <div className="my-4">
            <h1 className="mb-2 text-2xl font-bold">Itinerary</h1>

            <Itinerary  data={[]} />
          </div>
        </div>
      </div>
    </Scroll>
  )
}

export default Match