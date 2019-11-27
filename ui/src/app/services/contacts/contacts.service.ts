import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {Contact} from '../../model/contact';
import {GenericDataService} from '../GenericDataService';

@Injectable({
  providedIn: 'root',
})
export class ContactsService extends GenericDataService {
  private readonly getAllContactsURL = this.server + '/contacts';

  constructor(private http: HttpClient) { super(); }

  getAllContacts(): Observable<Contact[]> {
    return this.http.get<Contact[]>(this.getAllContactsURL);
  }
}
