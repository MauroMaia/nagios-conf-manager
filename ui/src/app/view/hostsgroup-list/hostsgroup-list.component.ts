import {NestedTreeControl} from '@angular/cdk/tree';
import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs';
import {HostGroup} from '../../model/host-group';

import {HostgroupDataService} from '../../services/hostgroup/hostgroup-data.service';

interface TreeNode
{
  name: string;

  children?: TreeNode[];
}

@Component({
  selector   : 'app-hostsgroup-list',
  templateUrl: './hostsgroup-list.component.html',
  styleUrls  : ['./hostsgroup-list.component.scss'],
})
export class HostsgroupListComponent implements OnInit {

  private hostGroupsSub: Observable<HostGroup[]>;

  // tslint:disable-next-line:no-shadowed-variable
  constructor(private HostDataService: HostgroupDataService) { }

  ngOnInit() {
    // @ts-ignore
    this.hostGroupsSub = this.HostDataService.getAllHostGroups();
  }
}

