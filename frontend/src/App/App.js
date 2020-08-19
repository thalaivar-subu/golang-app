import React from "react";
import WordCountForm from "../WordCountForm/WordCountForm";
import "./App.css";
import ExcelForm from "../ExcelForm/ExcelForm";
import CrudFrom from "../CrudForm/CrudFrom";
import PrimeForm from "../PrimeForm/PrimeForm";


function App() {
  return (
    <div className="App">
      <h1>Html Word Count</h1>
      <WordCountForm />
      <br />
      <h1>Excel Column Name</h1>
      <ExcelForm />
      <br />
      <h1>Find Prime Number</h1>
      <PrimeForm />
      <br />
      <h1>CRUD Operations</h1>
      <CrudFrom />
      <br />
    </div>
  );
}

export default App;
