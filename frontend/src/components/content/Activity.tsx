import "./Activity.css";
import Cookies from 'universal-cookie';

function Activity() {
  const cookies = new Cookies();
  console.log('From Activity: '+cookies.get('myCat')); // Pacman
  return <div>Activity</div>;
}

export default Activity;
