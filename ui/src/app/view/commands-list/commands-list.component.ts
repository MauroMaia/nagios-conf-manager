import { Component, OnInit } from '@angular/core';
import {Observable} from 'rxjs';
import {Command} from '../../model/command';
import {CommandsService} from '../../services/commands/commands.service';

@Component({
  selector: 'app-commands-list',
  templateUrl: './commands-list.component.html',
  styleUrls: ['./commands-list.component.scss']
})
export class CommandsListComponent implements OnInit {

  private commandsObservable: Observable<Command[]>;

  constructor(private commandsService: CommandsService)  { }

  ngOnInit() {
    this.commandsObservable = this.commandsService.getAllCommands();
    // @ts-ignore
    this.commandsObservable.subscribe(array => console.log(array));
  }

}
