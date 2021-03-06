import {Component, OnInit} from '@angular/core';
import {FormControl, Validators} from '@angular/forms';
import {of} from 'rxjs';

@Component({
  selector   : 'app-add-timeperiod',
  templateUrl: './add-timeperiod.component.html',
  styleUrls  : ['./add-timeperiod.component.scss'],
})
export class AddTimeperiodComponent implements OnInit
{
  // tslint:disable-next-line:variable-name
  public tpName = new FormControl('',
    [Validators.required, Validators.pattern(/^[a-zA-Z0-9|_|-]+$/)]);

  public tpDescription = new FormControl('',
    [Validators.required, Validators.pattern(/.*/)]);

  public errorsStrings = {
    name          : {
      required: 'You must enter a value',
      pattern : 'Not a valid pattern for name field',
    }, description: {
      required: 'You must enter a value',
      pattern : 'Not a valid pattern for description field',
    },
  };

  public disableStatus = {
    monday   : true,
    tuesday  : true,
    wednesday: true,
    thursday : true,
    friday   : true,
    saturday : true,
    sunday   : true,
  };

  public disableStatusOb = {
    monday   : of(this.disableStatus.monday),
    tuesday  : of(this.disableStatus.tuesday),
    wednesday: of(this.disableStatus.wednesday),
    thursday : of(this.disableStatus.thursday),
    friday   : of(this.disableStatus.friday),
    saturday : of(this.disableStatus.saturday),
    sunday   : of(this.disableStatus.sunday),
  };


  constructor() { }

  ngOnInit() { }

}
