import {ReactiveFormsModule} from '@angular/forms';
import {
  MatButtonModule, MatExpansionModule,
  MatIconModule,
  MatListModule,
  MatNativeDateModule,
  MatSidenavModule,
  MatTableModule,
  MatToolbarModule, MatTreeModule,
} from '@angular/material';
import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {AppComponent} from './app.component';
import {CommandsService} from './services/commands/commands.service';
import {ContactGroupsService} from './services/contactGroup/contact-groups.service';
import {ContactsService} from './services/contacts/contacts.service';
import {HostDataService} from './services/host/host-data.service';
import {HttpClientModule} from '@angular/common/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {HostgroupDataService} from './services/hostgroup/hostgroup-data.service';
import {ServiceService} from './services/service/service.service';
import {HostsgroupListComponent} from './view/hostsgroup-list/hostsgroup-list.component';
import {SidebarComponent} from './view/sidebar/sidebar.component';
import {HostsListComponent} from './view/hosts-list/hosts-list.component';
import {TimeperiodListComponent} from './view/timeperiod-list/timeperiod-list.component';
import {CommandsListComponent} from './view/commands-list/commands-list.component';
import {ServiceListComponent} from './view/service-list/service-list.component';
import {ContactListComponent} from './view/contact-list/contact-list.component';
import {ContactGroupListComponent} from './view/contact-group-list/contact-group-list.component';

const appRoutes: Routes = [
  /*  { path: 'crisis-center', component: CrisisListComponent },
    { path: 'hero/:id',      component: HeroDetailComponent },
    {
      path: 'heroes',
      component: HeroListComponent,
      data: { title: 'Heroes List' }
    },
    { path: '**', component: PageNotFoundComponent }*/
  {
    path      : '',
    redirectTo: '/hosts',
    pathMatch : 'full',
  }, {
    path     : 'hosts',
    pathMatch: 'full',
    component: HostsListComponent,
  }, {
    path     : 'hostsgroup',
    pathMatch: 'full',
    component: HostsgroupListComponent,
  }, {
    path     : 'timeperiods',
    pathMatch: 'full',
    component: TimeperiodListComponent,
  }, {
    path     : 'commands',
    pathMatch: 'full',
    component: CommandsListComponent,
  }, {
    path     : 'services',
    pathMatch: 'full',
    component: ServiceListComponent,
  }, {
    path     : 'contacts',
    pathMatch: 'full',
    component: ContactListComponent,
  }, {
    path     : 'contactgroups',
    pathMatch: 'full',
    component: ContactGroupListComponent,
  },
];

@NgModule({
  declarations: [
    AppComponent,
    SidebarComponent,
    HostsListComponent,
    HostsgroupListComponent,
    TimeperiodListComponent,
    CommandsListComponent,
    ServiceListComponent,
    ContactListComponent,
    ContactGroupListComponent,
  ],
  imports     : [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule, // import HttpClientModule after BrowserModule.
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    MatButtonModule,
    MatIconModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    RouterModule.forRoot(
      appRoutes,
      {enableTracing: true}, // <-- debugging purposes only
    ),
    MatTableModule,
    MatExpansionModule,
    MatTreeModule,
  ],
  providers   : [
    HostDataService,
    HostgroupDataService,
    CommandsService,
    ServiceService,
    ContactsService,
    ContactGroupsService,
  ],
  exports     : [],
  bootstrap   : [AppComponent],
})
export class AppModule {}
