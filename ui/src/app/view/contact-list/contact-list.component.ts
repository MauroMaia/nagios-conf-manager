import { Component, OnInit } from '@angular/core';
import {Observable} from 'rxjs';
import {Contact} from '../../model/contact';
import {ContactsService} from '../../services/contacts/contacts.service';


@Component({
  selector: 'app-contact-list',
  templateUrl: './contact-list.component.html',
  styleUrls: ['./contact-list.component.scss']
})
export class ContactListComponent implements OnInit {

  public contactObservableList: Observable<Contact[]>;
  // tslint:disable-next-line:no-shadowed-variable
  constructor(private contactsService: ContactsService)  { }

  ngOnInit() {
    this.contactObservableList = this.contactsService.getAllContacts();
    // @ts-ignore
    this.contactObservableList.subscribe(array => console.log(array));
  }

}
