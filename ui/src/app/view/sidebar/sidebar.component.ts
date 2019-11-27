import {Component, OnInit} from '@angular/core';

@Component({
  selector   : 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls  : ['./sidebar.component.scss'],
})
export class SidebarComponent implements OnInit {

  public readonly menus = [
    'hosts',
    'hostsgroup',
    'timeperiods',
    'commands',
    'services',
    'contacts',
    'contactgroups'
  ];

  constructor() { }

  ngOnInit() { }

}
