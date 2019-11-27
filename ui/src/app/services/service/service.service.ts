import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {Command} from '../../model/command';
import {Service} from '../../model/service';
import {GenericDataService} from '../GenericDataService';

@Injectable({
  providedIn: 'root',
})
export class ServiceService extends GenericDataService {

  private readonly getAllServicesURL = this.server + '/services';

  constructor(private http: HttpClient) { super(); }

  getAllServices(): Observable<Service[]> {
    return this.http.get<Service[]>(this.getAllServicesURL);
    /* .pipe(
       tap(_ => console.info('fetched hosts')),
       catchError(this.handleError<Host[]>('getHeroes', []))
     );*/
  }

}
