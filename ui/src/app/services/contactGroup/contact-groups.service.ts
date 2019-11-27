import {HttpClient} from '@angular/common/http';
import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {ContactGroup} from '../../model/contact-group';
import {GenericDataService} from '../GenericDataService';

@Injectable({
  providedIn: 'root'
})
export class ContactGroupsService extends GenericDataService {

  private readonly getAllContactGroupURL = this.server + '/contactgroups';

  constructor(private http: HttpClient) { super(); }

  getAllContactGroups(): Observable<ContactGroup[]> {
    return this.http.get<ContactGroup[]>(this.getAllContactGroupURL);
    /* .pipe(
       tap(_ => console.info('fetched hosts')),
       catchError(this.handleError<Host[]>('getHeroes', []))
     );*/
  }
}
