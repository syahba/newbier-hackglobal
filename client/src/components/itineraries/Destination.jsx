import CardAction from './CardAction';
import CardDefault from './CardDefault';

function Destination() {
  return (
    <div className="ml-2 border-l-2 border-blue">
      <div className="relative pl-2">
        <CardDefault></CardDefault>
        <p className="my-3 text-center">300 meters - 20 minutes</p>
      </div>
      <div className="relative pl-2">
        <CardAction></CardAction>
        <p className="my-3 text-center">300 meters - 20 minutes</p>
      </div>
    </div>
  );
}

export default Destination;
