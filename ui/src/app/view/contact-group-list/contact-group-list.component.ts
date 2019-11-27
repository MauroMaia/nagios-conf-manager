import { Component, OnInit } from '@angular/core';
import {Observable} from 'rxjs';
import {ContactGroup} from '../../model/contact-group';
import {ContactGroupsService} from '../../services/contactGroup/contact-groups.service';

@Component({
  selector: 'app-contact-group-list',
  templateUrl: './contact-group-list.component.html',
  styleUrls: ['./contact-group-list.component.scss']
})
export class ContactGroupListComponent implements OnInit {

  public contactGroupObservableList: Observable<ContactGroup[]>;

  // tslint:disable-next-line:no-shadowed-variable
  constructor(private contactGroupsService: ContactGroupsService)  { }

  ngOnInit() {
    this.contactGroupObservableList = this.contactGroupsService.getAllContactGroups();
    // @ts-ignore
    this.contactGroupObservableList.subscribe(array => console.log(array));
  }

}
