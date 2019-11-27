import { Component, OnInit } from '@angular/core';
import {Observable} from 'rxjs';
import {TimePeriod} from '../../model/time-periods';
import {TimePeriodService} from '../../services/timeperiod/time-period.service';

@Component({
  selector: 'app-timeperiod-list',
  templateUrl: './timeperiod-list.component.html',
  styleUrls: ['./timeperiod-list.component.scss']
})
export class TimeperiodListComponent implements OnInit {
  private timePeriodObservable: Observable<TimePeriod[]>;

  constructor(private timePeriodService: TimePeriodService)  { }

  ngOnInit() {
    this.timePeriodObservable = this.timePeriodService.getAllHosts();
    // @ts-ignore
    this.timePeriodObservable.subscribe(array => console.log(array));
  }

}
