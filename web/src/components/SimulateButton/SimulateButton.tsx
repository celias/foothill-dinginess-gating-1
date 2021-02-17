import React from 'react';
import Button from '@material-ui/core/Button';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import PlayCircleFilledIcon from '@material-ui/icons/PlayCircleFilled';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    button: {
      color: 'white',
      backgroundColor: '#7E57C2'
    },
  }),
);

export default function SimulateButton() {
  const classes = useStyles();

  return (
    <div className="SimulateButton">
      <Button
        variant="contained"
        className={classes.button}
        startIcon={<PlayCircleFilledIcon />}
        fullWidth
      >
        Simulate
      </Button>
    </div>
  );
}
