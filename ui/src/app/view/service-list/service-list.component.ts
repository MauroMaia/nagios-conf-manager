import { Component, OnInit } from '@angular/core';
import {Observable} from 'rxjs';
import {Service} from '../../model/service';
import {ServiceService} from '../../services/service/service.service';

@Component({
  selector: 'app-service-list',
  templateUrl: './service-list.component.html',
  styleUrls: ['./service-list.component.scss']
})
export class ServiceListComponent implements OnInit {

  private servicesListObservable: Observable<Service[]>;

  // tslint:disable-next-line:no-shadowed-variable
  constructor(private serviceService: ServiceService) { }

  ngOnInit() {
    // @ts-ignore
    this.servicesListObservable = this.serviceService.getAllServices();
  }

}
