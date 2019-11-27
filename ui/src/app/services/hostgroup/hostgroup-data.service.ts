import {HttpClient} from '@angular/common/http';
import { Injectable } from '@angular/core';
import {Observable, of} from 'rxjs';
import {HostGroup} from '../../model/host-group';

@Injectable({
  providedIn: 'root'
})
export class HostgroupDataService {

  server = 'http://127.0.0.1:3000/hostgroups';


  constructor(private http: HttpClient) {}

  getAllHostGroups(): Observable<HostGroup[]> {
    return this.http.get<HostGroup[]>(this.server);
    /*.pipe(
       tap(_ => console.info('fetched hosts')),
       catchError(this.handleError<Host[]>('getHeroes', []))
     );*/
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      console.error(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
