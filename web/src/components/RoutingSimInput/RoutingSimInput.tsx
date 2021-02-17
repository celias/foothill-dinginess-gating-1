import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import SimulateButton from '../SimulateButton/SimulateButton';

export default function RoutingSimInput() {

  return (
    <div className="RoutingSimInput">
    <form className="RoutingSimInput__form" noValidate autoComplete="off">
      <TextareaAutosize
        rowsMax={4}
        aria-label="text area"
      />
    </form>

    <SimulateButton />

  </div>
   
  );
}


