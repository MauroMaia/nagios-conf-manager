import { Component, OnInit } from '@angular/core';
import {Observable} from 'rxjs';
import {Host} from '../../model/host';
import {HostDataService} from '../../services/host/host-data.service';

@Component({
  selector: 'app-hosts-list',
  templateUrl: './hosts-list.component.html',
  styleUrls: ['./hosts-list.component.scss']
})
export class HostsListComponent implements OnInit {
  public hostSubs: Observable<Host[]>;

  // tslint:disable-next-line:no-shadowed-variable
  constructor(private HostDataService: HostDataService)  { }

  ngOnInit() {
    this.hostSubs = this.HostDataService.getAllHosts();
    // @ts-ignore
    this.hostSubs.subscribe(array => console.log(array));
  }

}
