import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {FormControl, Validators} from '@angular/forms';
import {Observable} from 'rxjs';

@Component({
  selector   : 'app-time-period-day-ui',
  templateUrl: './time-period-da-yui.component.html',
  styleUrls  : ['./time-period-da-yui.component.scss'],
})
export class TimePeriodDaYUIComponent
{

  @Input() disable: Observable<boolean>;

  @Output() disableChange: EventEmitter<boolean> = new EventEmitter<boolean>();


  @Input() placeholder = '';

  private pattern = /^[0-9]{1,2}:[0-9]{2}-[0-9]{1,2}:[0-9]{2}$/;

  public tpDayOfWeek = new FormControl('',
    [Validators.pattern(this.pattern)]);

  public errorsStrings = {
    dayOfWeek: {
      required: 'You must enter a value',
      pattern : 'Not a valid pattern. Must match: ' + this.pattern,
    },
  };

  constructor() { }

  public changeStatus()
  {
    this.disableChange.emit(!this.disable);
  }

}
