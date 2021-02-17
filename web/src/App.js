import HorizontalDivider from "../src/components/Dividers/HorizontalDivider/HorizontalDivider";
import RoutingSimInput from "./components/RoutingSimInput/RoutingSimInput";
import "./App.css";

function App() {
  return (
    <div className='App'>
      <div className='App__container'>
        <div className='App__RouteSim'>
          <header className='App__header'>
            <h2>Routing Sim 1000</h2>
          </header>
          <label>Input CSV Data</label>
          <RoutingSimInput />
        </div>

        <div className='App__Instructions'>
          <label>Instructions</label>
          {/* <div className='App__Instructions-list'> */}
          <ol>
            <li>Provide your input data</li>
            <li>Click Simulate</li>
            <li>Review Results</li>
          </ol>
          {/* </div> */}
          <hr />
        </div>
      </div>
      <HorizontalDivider />
    </div>
  );
}

export default App;
