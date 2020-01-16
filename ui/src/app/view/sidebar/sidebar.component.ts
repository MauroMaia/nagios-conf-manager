import {Component, OnInit} from '@angular/core';
import {MatIconRegistry} from '@angular/material';
import {DomSanitizer} from '@angular/platform-browser';

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

  constructor(iconRegistry: MatIconRegistry, sanitizer: DomSanitizer) {
    iconRegistry.addSvgIcon(
      'add-24px',
      sanitizer.bypassSecurityTrustResourceUrl('assets/add-24px.svg'),
    );
  }

  ngOnInit() { }

}
