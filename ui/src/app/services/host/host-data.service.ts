import {Injectable} from '@angular/core';
import {Host} from '../../model/host';
import {Observable, of} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import {GenericDataService} from '../GenericDataService';


@Injectable({
  providedIn: 'root',
})
export class HostDataService extends GenericDataService {

  private readonly getAllHostsUrl = this.server + '/hosts';


  constructor(private http: HttpClient) {
    super();
  }

  getAllHosts(): Observable<Host[]> {
    return this.http.get<Host[]>(this.getAllHostsUrl);
    /* .pipe(
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
